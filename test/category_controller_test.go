package test

import (
	"database/sql"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
 
func setupRouter() http.Handler	{
	db := setupTestDB()
	validate := validator.New(validator.WithRequiredStructEnabled())

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)


	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)

}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"name" : "perabotan"}`)
	request := httptest.NewRequest("POST", "http://localhost:8000/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "password")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestCreateCategoryFailed(t *testing.T) {

}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestGetAllCategorySuccess(t *testing.T) {

}

func TestGetAllCategoryFailed(t *testing.T) {

}

func TestAuthorizedSuccess(t *testing.T) {

}

func TestAuthorizedFailed(t *testing.T) {

}