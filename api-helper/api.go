package apihelper

import (
	"golang-restful-api/middleware"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func SecureApi(middleware middleware.ApiKeyAuthMiddleware, handler httprouter.Handle, except ...string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for _, e := range except {
			if r.URL.Path == e {
				handler(w,r,p)
				return
			}
		}
		middleware.Wrap(handler)(w, r, p)
	}
}
