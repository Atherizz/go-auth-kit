package service

import (
	"context"
	"golang-restful-api/model/web"
)

type AuthService interface {
	CheckCredentials(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error)
	GetById(ctx context.Context, id int) (web.UserResponse, error)
	GetByColumn(ctx context.Context, data string, column string) (web.UserResponse, error)
	Register(ctx context.Context, request web.UserRequest) web.UserResponse
	SetVerified(ctx context.Context, token string) (web.UserResponse,error)
	ResendVerifyToken(ctx context.Context, email string) (web.VerifyTokenResponse,error)
	ForgotPassword(ctx context.Context, email string) (web.ResetTokenResponse,error)
	ResetPassword(ctx context.Context, request web.ResetPasswordRequest, token string) error
	ChangePassword(ctx context.Context, request web.ResetPasswordRequest, id int) error
	
}
