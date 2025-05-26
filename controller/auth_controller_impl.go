package controller

import (
	"encoding/json"
	"fmt"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/service"
	"golang-restful-api/model/web"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(AuthService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: AuthService,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)
	userCreateRequest := web.UserRequest{}
	err := decoder.Decode(&userCreateRequest)
	helper.PanicError(err)

	userResponse := controller.AuthService.Register(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)

}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := web.LoginRequest{}
	err := decoder.Decode(&loginRequest)
	helper.PanicError(err)

	response, err := controller.AuthService.CheckCredentials(request.Context(), loginRequest)
	// fmt.Fprint(writer, response.Data)

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

func (controller *AuthControllerImpl) CheckUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId")
	if userId == nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := userId.(int)
	fmt.Printf("User ID from context: %d\n", userID)
	hello := "hello user " + strconv.Itoa(userID)

	helper.WriteEncodeResponse(writer, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   hello,
	})
}

func (controller *AuthControllerImpl) GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId")
	if userId == nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := userId.(int)

	response, err := controller.AuthService.GetById(request.Context(), userID)

	if err != nil {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteEncodeResponse(writer, webResponse)

}

func (controller *AuthControllerImpl) VerifyUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keyword := request.URL.Query().Get("token")

	user, err := controller.AuthService.GetByColumn(request.Context(), keyword, "verify_token")
	token := user.VerifyToken

	if err != nil {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	if keyword != token {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Token false!",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	if time.Now().After(user.ExpiredAt) {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Token already expired!",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	if user.IsVerify == 1 {
		webResponse := web.WebResponse{
			Code:   200,
			Status: "User already verified",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	verifyUser, err := controller.AuthService.SetVerified(request.Context(), token)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Verify Success",
		Data:   verifyUser,
	}
	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *AuthControllerImpl) ResendVerifyToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	emailRequest := web.EmailRequest{}
	err := decoder.Decode(&emailRequest)
	helper.PanicError(err)

	response, err := controller.AuthService.ResendVerifyToken(request.Context(), emailRequest.Email)
	if err != nil {
		webResponse := web.WebResponse{
		Code:   400,
		Status: "Email Not Found / Not registered",
		Data:   nil,
	}
	helper.WriteEncodeResponse(writer, webResponse)
	return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Resend Verification token success",
		Data:   response,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *AuthControllerImpl) ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	emailRequest := web.EmailRequest{}
	err := decoder.Decode(&emailRequest)
	helper.PanicError(err)

	response, err := controller.AuthService.ForgotPassword(request.Context(), emailRequest.Email)
	if err != nil {
		webResponse := web.WebResponse{
		Code:   400,
		Status: "Email Not Found / Not registered",
		Data:   nil,
	}
	helper.WriteEncodeResponse(writer, webResponse)
	return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Send Reset token success",
		Data:   response,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *AuthControllerImpl) ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	resetPasswordRequest := web.ResetPasswordRequest{}
	err := decoder.Decode(&resetPasswordRequest)
	helper.PanicError(err)

	token := request.URL.Query().Get("token")

	user, err := controller.AuthService.GetByColumn(request.Context(), token, "reset_token")
	if err != nil {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	if time.Now().After(user.ResetExpiredAt) {
		webResponse := web.WebResponse{
			Code:   401,
			Status: "Token already expired!",
			Data:   nil,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return
	}

	err = controller.AuthService.ResetPassword(request.Context(), resetPasswordRequest, token)
	if err != nil {
		webResponse := web.WebResponse{
		Code:   400,
		Status: "Reset Password Failed!",
		Data:   nil,
	}
	helper.WriteEncodeResponse(writer, webResponse)
	return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Reset Password Success!",
	}

	helper.WriteEncodeResponse(writer, webResponse)
}
