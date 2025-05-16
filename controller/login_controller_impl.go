package controller

import (
	"encoding/json"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/service"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LoginControllerImpl struct {
	LoginService service.LoginService
}

func NewLoginController(loginService service.LoginService) *LoginControllerImpl {
	return &LoginControllerImpl{
		LoginService: loginService,
	}
}

func (controller *LoginControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := web.LoginRequest{}
	err := decoder.Decode(&loginRequest)
	helper.PanicError(err)

	response, err := controller.LoginService.CheckCredentials(request.Context(), loginRequest)
	
	if err != nil {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Unauthorized - Invalid email or password",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Login Success",
		Data:   response,
	}

	helper.WriteEncodeResponse(writer, webResponse)

}
