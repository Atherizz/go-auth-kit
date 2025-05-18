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

type RecipeServiceImpl struct {
	RecipeRepository repository.RecipeRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewCategoryService(RecipeRepository repository.RecipeRepository, db *sql.DB, validate *validator.Validate) *RecipeServiceImpl {
	return &RecipeServiceImpl{
		RecipeRepository: RecipeRepository,
		DB:               db,
		Validate:         validate,
	}
}

func (service *RecipeServiceImpl) Create(ctx context.Context, request web.RecipeRequest) web.RecipeResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	recipe := entity.Recipe{
		Title:       request.Title,
		Ingredients: request.Ingredients,
		Calories:    float64(request.Calories),
		UserId:      request.UserId,
		CategoryId:  request.CategoryId,
	}

	recipe = service.RecipeRepository.Create(ctx, tx, recipe)

	return helper.ToRecipeResponse(recipe)

}

func (service *RecipeServiceImpl) Search(ctx context.Context, keyword string) []web.RecipeResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := service.RecipeRepository.Search(ctx, tx, keyword)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var recipesResponse []web.RecipeResponse

	for _, recipe := range categories {
		recipesResponse = append(recipesResponse, helper.ToRecipeResponse(recipe))
	}

	return recipesResponse

}

func (service *RecipeServiceImpl) FindById(ctx context.Context, id int) web.RecipeResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	recipe, err := service.RecipeRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRecipeResponse(recipe)

}

func (service *RecipeServiceImpl) Show(ctx context.Context) []web.RecipeResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.RecipeRepository.GetAll(ctx, tx)
	var recipesResponse []web.RecipeResponse

	for _, category := range categories {
		recipesResponse = append(recipesResponse, helper.ToRecipeResponse(category))
	}

	return recipesResponse

}

func (service *RecipeServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.RecipeRepository.GetById(ctx, tx, id)
	helper.PanicError(err)

	err = service.RecipeRepository.Delete(ctx, tx, int32(id))
	if err != nil {
		return err
	}

	return nil

}
