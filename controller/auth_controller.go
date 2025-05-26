package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CheckUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	VerifyUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ResendVerifyToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}