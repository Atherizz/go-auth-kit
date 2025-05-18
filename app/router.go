package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/middleware"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryControllers, userControllers controller.EntityController[web.EntityRequest, entity.NamedEntity, web.EntityResponse], loginController controller.LoginController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", userControllers.Create)
	router.POST("/api/login", loginController.Login)

	jwtMiddleware := middleware.NewJwtAuthMiddleware(router)

	router.GET("/api/categories", jwtMiddleware.Wrap(categoryControllers.FindAll))
	router.GET("/api/categories/:entityId", jwtMiddleware.Wrap(categoryControllers.FindById))
	router.POST("/api/categories", jwtMiddleware.Wrap(categoryControllers.Create))
	router.PUT("/api/categories/:entityId", jwtMiddleware.Wrap(categoryControllers.Update))
	router.DELETE("/api/categories/:entityId", jwtMiddleware.Wrap(categoryControllers.Delete))

	router.GET("/api/users", jwtMiddleware.Wrap(userControllers.FindAll))
	router.GET("/api/users/:entityId", jwtMiddleware.Wrap(userControllers.FindById))
	router.PUT("/api/users/:entityId", jwtMiddleware.Wrap(userControllers.Update))
	router.DELETE("/api/users/:entityId", jwtMiddleware.Wrap(userControllers.Delete))

	router.GET("/api/check-user", jwtMiddleware.Wrap(middleware.CheckUser))

	router.PanicHandler = exception.ErrorHandler
	router.MethodNotAllowed = http.HandlerFunc(exception.NotAllowedError)
	router.NotFound = http.HandlerFunc(exception.NotFoundRouteError)

	return router

}
