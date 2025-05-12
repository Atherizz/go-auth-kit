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

func (service *LoginServiceImpl) CheckCredentials(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.GetByName(ctx, tx, request.Email)
	userResponse := helper.ToUserResponse(user)
	helper.PanicError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicError(err)

	return web.LoginResponse{
		Status: "Success",
		Message: "Login success",
		Data: userResponse,
	}

	

}
