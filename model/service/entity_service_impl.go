package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exception"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"
	"log"

	"github.com/go-playground/validator/v10"
)

type ServiceImpl[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse] struct {
	Repository          repository.Repository[S]
	DB                  *sql.DB
	Validate            *validator.Validate
	ResponseConstructor func() R
}

func NewService[T web.EntityRequest, S entity.NamedEntity, R web.EntityResponse](categoryRepository repository.Repository[S], db *sql.DB, validate *validator.Validate, constructor func() R) *ServiceImpl[T, S, R] {
	return &ServiceImpl[T, S, R]{
		Repository:          categoryRepository,
		DB:                  db,
		Validate:            validate,
		ResponseConstructor: constructor,
	}
}

func (service *ServiceImpl[T, S, R]) Create(ctx context.Context, request T, model S) R {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	log.Printf("Creating: %+v\n", request)
	model.SetName(request.GetName())

	if model.GetEntityName() == "users" {
		model.SetEmail(request.GetEmail())
		model.SetPassword(request.GetPassword())
	}

	modelResult := service.Repository.Create(ctx, tx, model)
	log.Printf("Created model: %+v\n", modelResult)

	result := helper.ToCategoryResponse[S, R](modelResult, service.ResponseConstructor)
	log.Printf("Converted response: %+v\n", result)

	return result
}

func (service *ServiceImpl[T, S, R]) Update(ctx context.Context, request T, model S) R {

	errValidate := service.Validate.Var(request.GetName(), "required,min=1,max=200")
	helper.PanicError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	existingModel, err := service.Repository.GetById(ctx, tx, request.GetId(), model)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	existingModel.SetName(request.GetName())
	// existingModel.SetId(request.GetId())

	modelResult, err := service.Repository.Update(ctx, tx, existingModel)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := helper.ToCategoryResponse[S, R](modelResult, service.ResponseConstructor)
	return result
}

func (service *ServiceImpl[T, S, R]) FindById(ctx context.Context, id int, request T, model S) R {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	modelResult, err := service.Repository.GetById(ctx, tx, id, model)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := helper.ToCategoryResponse[S, R](modelResult, service.ResponseConstructor)
	return result
}

func (service *ServiceImpl[T, S, R]) Search(ctx context.Context, keyword string, request T, model S) []R {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	datas, err := service.Repository.Search(ctx, tx, keyword, model)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var categoriesResponse []R

	for _, data := range datas {
		categoryResponse := helper.ToCategoryResponse[S, R](data, service.ResponseConstructor)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}

func (service *ServiceImpl[T, S, R]) Show(ctx context.Context, request T, model S) []R {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	datas := service.Repository.GetAll(ctx, tx, model)
	var categoriesResponse []R

	for _, data := range datas {
		categoryResponse := helper.ToCategoryResponse[S, R](data, service.ResponseConstructor)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse

}

func (service *ServiceImpl[T, S, R]) Delete(ctx context.Context, id int, model S) error {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.Repository.GetById(ctx, tx, id, model)
	helper.PanicError(err)

	err = service.Repository.Delete(ctx, tx, int32(id), model)
	if err != nil {
		return err
	}

	return nil

}
