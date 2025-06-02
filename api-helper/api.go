package apihelper

import (
	"golang-restful-api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SecureRoute(middleware middleware.ApiKeyAuthMiddleware, router *httprouter.Router, except ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, e := range except {
			if r.URL.Path == e {
				router.ServeHTTP(w, r)
				return
			}
		}
		middleware.WrapRouter(router).ServeHTTP(w, r)
	})
}
