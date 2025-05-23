package controller

import (
	"encoding/json"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/service"
	"golang-restful-api/model/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RecipeControllerImpl struct {
	RecipeService service.RecipeService
}

func NewRecipeController(RecipeService service.RecipeService) *RecipeControllerImpl {
	return &RecipeControllerImpl{
		RecipeService: RecipeService,
	}
}

func (controller *RecipeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)
	recipeCreateRequest := web.RecipeRequest{}
	err := decoder.Decode(&recipeCreateRequest)
	helper.PanicError(err)

	recipeResponse := controller.RecipeService.Create(request.Context(), recipeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   recipeResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)

}


func (controller *RecipeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	recipeId := params.ByName("recipeId")
	id, err := strconv.Atoi(recipeId)
	helper.PanicError(err)

	controller.RecipeService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteEncodeResponse(writer, webResponse)
}

func (controller *RecipeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	recipeId := params.ByName("recipeId")
	id, err := strconv.Atoi(recipeId)
	helper.PanicError(err)

	recipeResponse := controller.RecipeService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   recipeResponse,
	}

	helper.WriteEncodeResponse(writer, webResponse)
}


func (controller *RecipeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	keyword := request.URL.Query().Get("search")

	if keyword ==  "" {
		recipeResponses := controller.RecipeService.Show(request.Context())
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   recipeResponses,
		}
	
		helper.WriteEncodeResponse(writer, webResponse)
	} else {
		recipeResponses := controller.RecipeService.Search(request.Context(), keyword)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   recipeResponses,
		}
	
		helper.WriteEncodeResponse(writer, webResponse)
	}

}
