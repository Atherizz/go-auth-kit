package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/service"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New(validator.WithRequiredStructEnabled())
	response := func() *web.CategoryResponse {
		return &web.CategoryResponse{}
	}
	categoryRepository := repository.NewRepository[*entity.Category]()
	categoryService := service.NewService[*web.CategoryUpdateRequest, *entity.Category, *web.CategoryResponse](categoryRepository, db, validate, response)
	categoryController := controller.NewController[*web.CategoryUpdateRequest, *entity.Category, *web.CategoryResponse](categoryService, &web.CategoryUpdateRequest{}, &entity.Category{
		Entity: "categories",
		Column: []string{"name"},
	})

	router := app.NewRouter(categoryController)
	middleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware,
	}
	err := server.ListenAndServe()
	helper.PanicError(err)
}

// server := google_wire.InitializeServer()
