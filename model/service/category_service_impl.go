package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exception"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"golang-restful-api/model/repository"
	"golang-restful-api/model/web"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Update(ctx context.Context, id int, name string) web.CategoryResponse {

	errValidate := service.Validate.Var(name, "required,min=1,max=200")
	helper.PanicError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.CategoryRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category, err := service.CategoryRepository.Update(ctx, tx, id, name)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, id int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Search(ctx context.Context, keyword string) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := service.CategoryRepository.Search(ctx, tx, keyword)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	var categoriesResponse []web.CategoryResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, helper.ToCategoryResponse(category))
	}

	return categoriesResponse


}

func (service *CategoryServiceImpl) Show(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.GetAll(ctx, tx)
	var categoriesResponse []web.CategoryResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, helper.ToCategoryResponse(category))
	}

	return categoriesResponse

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.CategoryRepository.GetById(ctx, tx, id)
	helper.PanicError(err)

	err = service.CategoryRepository.Delete(ctx, tx, int32(id))
	if err != nil {
		return err
	}

	return nil

}
