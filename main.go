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
	categoryResponse := func() *web.CategoryResponse {
		return &web.CategoryResponse{}
	}

	userResponse := func() *web.UserResponse {
		return &web.UserResponse{}
	}
	categoryRepository := repository.NewRepository[*entity.Category]()
	categoryService := service.NewService[*web.CategoryRequest, *entity.Category, *web.CategoryResponse](categoryRepository, db, validate, categoryResponse)
	categoryController := controller.NewController[*web.CategoryRequest, *entity.Category, *web.CategoryResponse](categoryService, &web.CategoryRequest{}, &entity.Category{
		Column: []string{"name"},
	})

	userRepository := repository.NewRepository[*entity.User]()
	userService := service.NewService[*web.UserRequest, *entity.User, *web.UserResponse](userRepository, db, validate, userResponse)
	userController := controller.NewController[*web.UserRequest, *entity.User, *web.UserResponse](userService, &web.UserRequest{}, &entity.User{
		Column: []string{"name", "email", "password_hash"},
	})

	loginRepository := repository.NewAuthRepository()
	loginService := service.NewAuthService(loginRepository, db, validate)
	loginController := controller.NewAuthController(loginService)

	recipeRepository := repository.NewRecipeRepository()
	recipeService := service.NewRecipeService(recipeRepository,db,validate)
	recipeController := controller.NewRecipeController(recipeService)

	router := app.NewRouter(categoryController, userController, loginController, recipeController, db)
	apiKeyMiddleware := middleware.NewApiKeyAuthMiddleware(router)
	
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: apiKeyMiddleware,
	}
	err := server.ListenAndServe()
	helper.PanicError(err)
}

// server := google_wire.InitializeServer()
