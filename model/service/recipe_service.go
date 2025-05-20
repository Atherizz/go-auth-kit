package service

import (
	"context"
	"golang-restful-api/model/web"
)

type RecipeService interface {
	
	Create(ctx context.Context, request web.RecipeRequest) web.RecipeResponse
	FindById(ctx context.Context, id int) web.RecipeResponse
	Show(ctx context.Context) []web.RecipeResponse
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, keyword string) []web.RecipeResponse
}