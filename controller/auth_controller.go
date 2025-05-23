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
}