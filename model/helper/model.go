package helper

import (
	"fmt"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
)

func ToCategoryResponse[S entity.NamedEntity, R web.EntityResponse](model S, constructor func() R) R {
	response := constructor()
	response.SetId(model.GetId())
	response.SetName(model.GetName())

	if model.GetEntityName() == "users" {
		fmt.Printf("DEBUG email: \n", model.GetEmail())
		fmt.Printf("DEBUG password: \n", model.GetPassword())
		response.SetEmail(model.GetEmail())
		response.SetPassword(model.GetPassword())
	}

	return response

}
