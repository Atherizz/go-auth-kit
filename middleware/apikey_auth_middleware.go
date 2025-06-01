package middleware

import (
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"net/http"
	"strings"

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

func (middleware ApiKeyAuthMiddleware) Wrap(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := "password"
	if strings.HasPrefix(request.URL.Path, "/api/") {
	if key != request.Header.Get("X-API-KEY") {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteEncodeResponse(writer, webResponse)
	} 
	
	next(writer,request,params)
	}

}
}

// func RecoveryMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
// 				log.Printf("Panic recovered: %v\n", err)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }