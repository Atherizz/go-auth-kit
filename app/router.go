package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryControllers, userControllers controller.Controller[web.EntityRequest, entity.NamedEntity, web.EntityResponse]) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryControllers.FindAll)
	router.GET("/api/categories/:entityId", categoryControllers.FindById)
	router.POST("/api/categories", categoryControllers.Create)
	router.PUT("/api/categories/:entityId", categoryControllers.Update)
	router.DELETE("/api/categories/:entityId", categoryControllers.Delete)

	router.POST("/api/register", userControllers.Create)
	router.GET("/api/users", userControllers.FindAll)
	router.GET("/api/users/:entityId", userControllers.FindById)
	router.PUT("/api/users/:entityId", userControllers.Update)
	router.DELETE("/api/users/:entityId", userControllers.Delete)

	router.PanicHandler = exception.ErrorHandler
	router.MethodNotAllowed = http.HandlerFunc(exception.NotAllowedError)
	router.NotFound = http.HandlerFunc(exception.NotFoundRouteError)

	return router

}
