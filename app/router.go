package app

import (
	"database/sql"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/middleware"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryControllers, userControllers controller.EntityController[web.EntityRequest, entity.NamedEntity, web.EntityResponse], authController controller.AuthController, recipeControllers controller.RecipeController, db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.GET("/api/verify-email", authController.VerifyUser)
	router.POST("/api/login", authController.Login)

	jwtMiddleware := middleware.NewJwtAuthMiddleware(router)
	adminMiddleware := middleware.NewAdminAuthMiddleware(router,db)
	checkUserMiddleware := middleware.NewCheckUserMiddleware(router)

	router.GET("/api/categories", jwtMiddleware.Wrap(categoryControllers.FindAll))
	router.GET("/api/categories/:entityId", jwtMiddleware.Wrap(categoryControllers.FindById))
	router.POST("/api/categories", jwtMiddleware.Wrap(adminMiddleware.Wrap(categoryControllers.Create)))
	router.PUT("/api/categories/:entityId", jwtMiddleware.Wrap(adminMiddleware.Wrap(categoryControllers.Update)))
	router.DELETE("/api/categories/:entityId", jwtMiddleware.Wrap(adminMiddleware.Wrap(categoryControllers.Delete)))

	router.GET("/api/users", jwtMiddleware.Wrap(userControllers.FindAll))
	router.GET("/api/users/:entityId", jwtMiddleware.Wrap(userControllers.FindById))
	router.PUT("/api/users/:entityId", jwtMiddleware.Wrap(checkUserMiddleware.Wrap(userControllers.Update)))
	router.DELETE("/api/users/:entityId", jwtMiddleware.Wrap(checkUserMiddleware.Wrap(userControllers.Delete)))

	router.GET("/api/recipes", jwtMiddleware.Wrap(recipeControllers.FindAll))
	router.GET("/api/recipes/:recipeId", jwtMiddleware.Wrap(recipeControllers.FindById))
	router.POST("/api/recipes", jwtMiddleware.Wrap(recipeControllers.Create))
	router.DELETE("/api/recipes/:recipeId", jwtMiddleware.Wrap(checkUserMiddleware.Wrap(recipeControllers.Delete)))

	router.GET("/api/check-user", jwtMiddleware.Wrap(authController.CheckUser))
	router.GET("/api/profile", jwtMiddleware.Wrap(authController.GetProfile))

	router.PanicHandler = exception.ErrorHandler
	router.MethodNotAllowed = http.HandlerFunc(exception.NotAllowedError)
	router.NotFound = http.HandlerFunc(exception.NotFoundRouteError)

	return router

}
