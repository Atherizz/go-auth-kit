package controller

import (
	"encoding/json"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/service"
	"golang-restful-api/model/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type EntityControllerImpl[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse] struct {
	CategoryService service.EntityService[T,S,R]
	Request        T
	Model           S

}

func NewController[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse](categoryService service.EntityService[T,S,R], request T, model S) *EntityControllerImpl[T,S,R] {
	return &EntityControllerImpl[T,S,R]{
		CategoryService: categoryService,
		Request: request,
		Model: model,
	}
}

func (controller *EntityControllerImpl[T, S,R]) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	createRequest := controller.Request
	err := decoder.Decode(&createRequest)
	// fmt.Printf("DEBUG createRequest: %+v\n", createRequest)
	helper.PanicError(err)


	dataResponse := controller.CategoryService.Create(request.Context(), createRequest, controller.Model)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)

}

func (controller *EntityControllerImpl[T, S, R]) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)
	updateRequest := controller.Request
	err := decoder.Decode(&updateRequest)
	helper.PanicError(err)

	entityId := params.ByName("entityId")
	id, err := strconv.Atoi(entityId)
	helper.PanicError(err)
	updateRequest.SetId(id)

	dataResponse := controller.CategoryService.Update(request.Context(), updateRequest, controller.Model)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *EntityControllerImpl[T, S, R]) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	entityId := params.ByName("entityId")
	id, err := strconv.Atoi(entityId)
	helper.PanicError(err)

	controller.CategoryService.Delete(request.Context(), id, controller.Model)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *EntityControllerImpl[T, S, R]) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	entityId := params.ByName("entityId")
	id, err := strconv.Atoi(entityId)
	helper.PanicError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id, controller.Request, controller.Model)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *EntityControllerImpl[T, S,R]) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	keyword := request.URL.Query().Get("search")

	if keyword == "" {
		categoryResponses := controller.CategoryService.Show(request.Context(), controller.Request, controller.Model)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponses,
		}

		helper.WriteEncodeResponse(writer, webResponse)
	} else {
		categoryResponses := controller.CategoryService.Search(request.Context(), keyword, controller.Request, controller.Model)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponses,
		}

		helper.WriteEncodeResponse(writer, webResponse)
	}

}
