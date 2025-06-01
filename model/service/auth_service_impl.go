package service

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"
	"log"
	"net/smtp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	Repository repository.AuthRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewAuthService(repo repository.AuthRepository, db *sql.DB, validate *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{
		Repository: repo,
		DB:         db,
		Validate:   validate,
	}
}
func (service *AuthServiceImpl) CheckCredentials(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.LoginResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.GetByColumn(ctx, tx, request.Email, "email")
	if err != nil {
		return web.LoginResponse{}, err
	}
	userResponse := helper.ToUserResponse(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.LoginResponse{}, err
	}

	secretKey := helper.LoadEnv("JWT_SECRET")

	token, err := helper.GenerateLoginToken(user.Id, user.Email, secretKey)
	if err != nil {
		return web.LoginResponse{}, err
	}

	loginResponse := web.LoginResponse{
		Data:  userResponse,
		Token: token,
	}

	return loginResponse, nil

}

func (service *AuthServiceImpl) GetByColumn(ctx context.Context, data string, column string) (web.UserResponse, error) {
	err := service.Validate.Var(data, "required,min=1")
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.GetByColumn(ctx, tx, data, column)
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse, nil
}

func (service *AuthServiceImpl) GetById(ctx context.Context, id int) (web.UserResponse, error) {
	err := service.Validate.Var(id, "number")
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.GetById(ctx, tx, id)
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse, nil
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	token := uuid.NewString()

	user := entity.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		VerifyToken: token,
		ExpiredAt:   time.Now().Add(15 * time.Minute),
	}

	user = service.Repository.Register(ctx, tx, user)

	auth := smtp.PlainAuth(
		"",
		"saveroathallah86@gmail.com",
		"qigkqalqvyaejjxk",
		"smtp.gmail.com",
	)

	msg := "Click here to verify your account\n" + "http://localhost:8000/api/verify-email?token="+token
	smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"saveroathallah86@gmail.com",
		[]string{request.Email},
		[]byte(msg),
	)

	return helper.ToUserResponse(user)

}

func (service *AuthServiceImpl) SetVerified(ctx context.Context, token string) (web.UserResponse, error) {
	err := service.Validate.Var(token, "required,min=1")
	if err != nil {
		log.Println("Validation error:", err)
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.SetVerified(ctx, tx, token)
	if err != nil {
		log.Println("Error in repository SetVerified:", err)
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}

func (service *AuthServiceImpl) ResendVerifyToken(ctx context.Context, email string) (web.VerifyTokenResponse, error) {
	err := service.Validate.Var(email, "required,email")
	if err != nil {
		return web.VerifyTokenResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.VerifyTokenResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.ResendVerifyToken(ctx, tx, email)
	if err != nil {
		return web.VerifyTokenResponse{}, err
	}

	return helper.ToVerifyTokenResponse(user), nil
}

func (service *AuthServiceImpl) ForgotPassword(ctx context.Context, email string) (web.ResetTokenResponse, error) {
	err := service.Validate.Var(email, "required,email")
	if err != nil {
		return web.ResetTokenResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return web.ResetTokenResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.ForgotPassword(ctx, tx, email)
	if err != nil {
		return web.ResetTokenResponse{}, err
	}
	return helper.ToResetTokenResponse(user), nil
}

func (service *AuthServiceImpl) ResetPassword(ctx context.Context, request web.ResetPasswordRequest, token string) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	err = service.Validate.Var(token, "required")
	if err != nil {
		return err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.Repository.ResetPassword(ctx, tx, request.NewPassword, token)
	if err != nil {
		return err
	}

	return nil

}

func (service *AuthServiceImpl) ChangePassword(ctx context.Context, request web.ResetPasswordRequest, id int) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	err = service.Validate.Var(id, "required,number")
	if err != nil {
		return err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.Repository.ChangePassword(ctx, tx, request.NewPassword, id)
	if err != nil {
		return err
	}

	return nil
}
