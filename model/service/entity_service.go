package service

import (
	"context"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	// "golang-restful-api/model/web"
)

type EntityService[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse] interface {
	Create(ctx context.Context, request T, model S) R
	Update(ctx context.Context, request T, model S) R
	FindById(ctx context.Context, id int, request T, model S) R
	Show(ctx context.Context, request T, model S) []R
	Delete(ctx context.Context, id int, model S) error
	Search(ctx context.Context, keyword string, request T, model S) []R
}
