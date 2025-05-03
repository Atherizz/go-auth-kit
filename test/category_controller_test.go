package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type data struct {
	Id   int
	Name string
}

type responseBody struct {
	Code   int
	Status string
	Data   data
}

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_migrations")
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New(validator.WithRequiredStructEnabled())

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)

}

func truncateData(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateData(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : "perabotan"}`)
	request := httptest.NewRequest("POST", "http://localhost:8000/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "password")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "perabotan", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()

	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest("POST", "http://localhost:8000/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "password")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	var res responseBody

	json.Unmarshal(body, &res)
	fmt.Println(res)

	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "BAD REQUEST", res.Status)
	assert.Equal(t, "Key: 'CategoryCreateRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag", res.Data)

}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateData(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{
	"name" : "gadget"
	}`)

	repo := repository.NewCategoryRepository()
	ctx := context.Background()

	tx, _ := db.Begin()
	newCategory := repo.Create(ctx, tx, entity.Category{
		Name: "aksesoris",
	})
	tx.Commit()

	request := httptest.NewRequest("PUT", "http://localhost:8000/api/categories/"+strconv.Itoa(newCategory.Id), requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "password")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var res responseBody
	json.Unmarshal(body, &res)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "OK", res.Status)
	assert.Equal(t, newCategory.Id, res.Data.Id)
	assert.Equal(t, "gadget", res.Data.Name)

}

func TestUpdateIdNotFound(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{
	"name" : "gadget"
	}`)

	request := httptest.NewRequest("PUT", "http://localhost:8000/api/categories/100", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "password")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var res responseBody
	json.Unmarshal(body, &res)

	assert.Equal(t, 404, res.Code)
	assert.Equal(t, "DATA NOT FOUND", res.Status)

}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateData(db)
	router := setupRouter(db)

	ctx := context.Background()
	tx, _ := db.Begin()
	repo := repository.NewCategoryRepository()
	category := repo.Create(ctx, tx, entity.Category{
		Name: "Makanan",
	})
	tx.Commit()

	request := httptest.NewRequest("GET", "http://localhost:8000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-KEY", "password")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var res responseBody
	json.Unmarshal(body, &res)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "OK", res.Status)
	assert.Equal(t, "Makanan", res.Data.Name)

}

func TestGetMethodNotAllowed(t *testing.T) {
	db := setupTestDB()
	truncateData(db)
	router := setupRouter(db)

	request := httptest.NewRequest("POST", "http://localhost:8000/api/categories/4", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-KEY", "password")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var res responseBody
	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
	assert.Equal(t, "METHOD NOT ALLOWED", res.Status)

}

func TestAuthorizedFailed(t *testing.T) {
	db := setupTestDB()
	truncateData(db)
	router := setupRouter(db)

	request := httptest.NewRequest("POST", "http://localhost:8000/api/categories/4", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var res responseBody
	json.Unmarshal(body, &res)

	assert.Equal(t, 401, res.Code)
	assert.Equal(t, "Unauthorized", res.Status)
	assert.Equal(t, "null", "null")
}
