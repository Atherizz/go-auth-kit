package controller

import (
	"fmt"
	"golang-restful-api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeControllerImpl struct {
}

func NewHomeController() *HomeControllerImpl {
	return &HomeControllerImpl{}
}

func (controller *HomeControllerImpl) BasicOauth(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "selamat datang di endpoint basic auth! anda berhasil terautentikasi \n")
}

func (controller *HomeControllerImpl) HomeOauth(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "welcome, you are authenticated \n")
}

func (controller *HomeControllerImpl) Callback(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	code := request.URL.Query().Get("code")
	token, err := middleware.OauthConfig.Exchange(request.Context(), code)
	if err != nil {
		http.Error(writer, "failed get token", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(writer, "you are authenticated! Token", token.AccessToken)
}
