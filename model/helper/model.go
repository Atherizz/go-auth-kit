package helper

import (
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
)

func ToCategoryResponse[S entity.NamedEntity, R web.EntityResponse](model S, constructor func() R) R {
	response := constructor()
	response.SetId(model.GetId())
	response.SetName(model.GetName())
	return response

}
