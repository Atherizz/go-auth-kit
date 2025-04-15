package app

import (
	"golang-restful-api/middleware"
	"net/http"
)

func NewServer(m *middleware.AuthMiddleware) *http.Server {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: m,
	}

	return &server
}
