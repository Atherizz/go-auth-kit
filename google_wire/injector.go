//go:build wireinject
// +build wireinject

package google_wire

import (
	"golang-restful-api/middleware"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/service"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeServer() (*http.Server) {
	wire.Build(
		app.NewServer,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewValidator,
		app.NewDB,
		categorySet,
	)

	return nil
}
