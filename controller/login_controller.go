package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type LoginController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CheckUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}