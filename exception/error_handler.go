package exception

import (
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func NotAllowedError(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)

	webResponse := web.WebResponse{
		Code:   http.StatusMethodNotAllowed,
		Status: "METHOD NOT ALLOWED",
	}
	helper.WriteEncodeResponse(writer, webResponse)
}

func NotFoundRouteError(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "UNKNOWN ENDPOINT",
	}
	helper.WriteEncodeResponse(writer, webResponse)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "DATA NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	var errorMsg string

	for _, e := range exception {
		if e.Tag() == "eqfield" {
			errorMsg = "Password confirmation does not match the password."
		} else if e.Tag() == "required" {
			errorMsg = "field is required"
		} else if e.Tag() == "email" {
			errorMsg = "field email must be an email"
		}
	}

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorMsg,
		}
		helper.WriteEncodeResponse(writer, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}
