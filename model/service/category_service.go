package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, id int, name string) web.CategoryResponse
	FindById(ctx context.Context, id int) web.CategoryResponse
	Show(ctx context.Context) []web.CategoryResponse
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, keyword string) []web.CategoryResponse
}
