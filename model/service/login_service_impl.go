package service

import (
	"context"
	"database/sql"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type LoginServiceImpl struct {
	Repository repository.LoginRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewLoginService(repo repository.LoginRepository, db *sql.DB, validate *validator.Validate) *LoginServiceImpl {
	return &LoginServiceImpl{
		Repository: repo,
		DB:         db,
		Validate:   validate,
	}
}
func (service *LoginServiceImpl) CheckCredentials(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.LoginResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.GetByEmail(ctx, tx, request.Email)
	if err != nil {
		return web.LoginResponse{}, err
	}
	userResponse := helper.ToUserResponse(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.LoginResponse{}, err
	}	

	secretKey := helper.LoadEnv("JWT_SECRET")

	token, err := helper.GenerateToken(user.Id, user.Email, secretKey)
	if err != nil {
		return web.LoginResponse{}, err
	}

	loginResponse := web.LoginResponse{
		Data:  userResponse,
		Token: token,
	}

	return loginResponse, nil

}
