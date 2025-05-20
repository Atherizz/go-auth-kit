package middleware

import (
	"database/sql"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AdminAuthMiddleware struct {
	Handler http.Handler
	DB      *sql.DB
}

func NewAdminAuthMiddleware(handler http.Handler, db *sql.DB) *AdminAuthMiddleware {
	return &AdminAuthMiddleware{
		Handler: handler,
		DB:      db,
	}
}

func (m *AdminAuthMiddleware) Wrap(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		userId := request.Context().Value("userId")
        log.Printf("DEBUG Middleware: userId = %v (type %T)", userId, userId)
        
		if userId == nil {
            helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}
        
		userRepo := repository.NewLoginRepository()

		tx, err := m.DB.Begin()
		if err != nil {
			helper.PanicError(err)
			return
		}
		defer tx.Rollback()
    

		user, err := userRepo.GetById(request.Context(), tx, userId.(int))
		helper.PanicError(err)
        log.Printf("DEBUG Middleware: user = %+v", user)

		if user.IsAdmin == 0 {
			helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}

		next(writer, request, param)

	}
}
