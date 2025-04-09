package helper

import (
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
)

func ToCategoryResponse(category entity.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
