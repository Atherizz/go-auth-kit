package middleware

import (
	"context"
	"database/sql"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type JwtAuthMiddlewareImpl struct {
	Handler http.Handler
	DB *sql.DB
}

func NewJwtAuthMiddleware(handler http.Handler, db *sql.DB) *JwtAuthMiddlewareImpl {
	return &JwtAuthMiddlewareImpl{
		Handler: handler,
		DB: db,
	}
}

func (m *JwtAuthMiddlewareImpl) Wrap(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		authHeader := request.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}


		secretKey := helper.LoadEnv("JWT_SECRET")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			helper.WriteEncodeResponse(writer, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
			})
			return
		}

		
		userID := int(claims["id"].(float64))

		userRepo := repository.NewAuthRepository()

        tx, err := m.DB.Begin()
        if err != nil {
            helper.PanicError(err)
            return
        }
        defer tx.Rollback()

        user, err := userRepo.GetById(request.Context(), tx, userID)
        helper.PanicError(err)

        if user.IsVerify == 0 {
                helper.WriteEncodeResponse(writer, web.WebResponse{
                Code:   http.StatusUnauthorized,
                Status: "Email Not Verified!",
            })
            return
        }

		ctx := context.WithValue(request.Context(), "userId", userID)
		ctx = context.WithValue(ctx, "email", claims["email"].(string))
		ctx = context.WithValue(ctx, "exp", claims["exp"].(float64))

		next(writer, request.WithContext(ctx), param)
	}
}