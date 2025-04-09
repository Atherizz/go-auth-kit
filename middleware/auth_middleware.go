package middleware

import (
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddlaware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddlaware {
	return &AuthMiddlaware{
		Handler: handler,
	}
}

func (middleware AuthMiddlaware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := "password"
	if key == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteEncodeResponse(writer, webResponse)

	}
}
