package main

import (
	"golang-restful-api/google_wire"
	"golang-restful-api/model/helper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := app.NewDB()
	// validate := validator.New(validator.WithRequiredStructEnabled())

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)


	// router := app.NewRouter(categoryController)
	// middleware := middleware.NewAuthMiddleware(router)

	// server := http.Server{
	// 	Addr:    "localhost:8000",
	// 	Handler: middleware,
	// }

	server := google_wire.InitializeServer()
	err := server.ListenAndServe()
	helper.PanicError(err)
}
