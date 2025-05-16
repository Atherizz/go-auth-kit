package service

import (
	"context"
	"golang-restful-api/model/web"
)

type LoginService interface {
	CheckCredentials(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error)
}
