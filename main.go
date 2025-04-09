package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New(validator.WithRequiredStructEnabled())

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)


	router := app.NewRouter(categoryController)


	middleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware,
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
