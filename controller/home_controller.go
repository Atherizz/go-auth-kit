package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeController interface {
	BasicOauth(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HomeOauth(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Callback(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
