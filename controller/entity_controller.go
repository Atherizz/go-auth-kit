package controller

import (
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type EntityController[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse] interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

}

