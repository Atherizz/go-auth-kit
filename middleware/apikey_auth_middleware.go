package middleware

import (
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ApiKeyAuthMiddleware struct {
	Handler http.Handler
}

func NewApiKeyAuthMiddleware(handler http.Handler) *ApiKeyAuthMiddleware {
	return &ApiKeyAuthMiddleware{
		Handler: handler,
	}
}

func (m *ApiKeyAuthMiddleware) WrapRouter(router *httprouter.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "password" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			}
			helper.WriteEncodeResponse(w, webResponse)
			return
		}
		router.ServeHTTP(w, r)
	})
}
