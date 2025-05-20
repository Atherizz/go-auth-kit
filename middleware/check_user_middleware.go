package middleware

import (
	"golang-restful-api/model/helper"
	"golang-restful-api/model/web"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CheckUserMiddleware struct {
	Handler http.Handler
}

func NewCheckUserMiddleware(handler http.Handler) *CheckUserMiddleware {
	return &CheckUserMiddleware{
		Handler: handler,
	}
}

func (m *CheckUserMiddleware) Wrap(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		userId := request.Context().Value("userId")

		entityId := params.ByName("entityId")
		id, _ := strconv.Atoi(entityId)
		
        log.Printf("DEBUG Middleware: userId = %v (type %T)", userId, userId)
        
		if userId == nil {
            helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}
    
		if id != userId.(int) {
			helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}

		next(writer, request, params)

	}
}
