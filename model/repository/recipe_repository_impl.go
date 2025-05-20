package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
)

type RecipeRepositoryImpl struct {
}

func NewRecipeRepository() *RecipeRepositoryImpl {
	return &RecipeRepositoryImpl{}

}

func (repo *RecipeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, recipe entity.Recipe) entity.Recipe {
	script := "INSERT INTO recipes(title, ingredients, calories, user_id, category_id) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, recipe.Title, recipe.Ingredients, recipe.Calories, recipe.UserId, recipe.CategoryId)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	recipe.Id = int(id)
	return recipe
}

func (repo *RecipeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Recipe {
	script := "SELECT id, title, ingredients, calories, user_id, category_id FROM recipes"

	result, err := tx.QueryContext(ctx, script)
	helper.PanicError(err)

	defer result.Close()

	var recipes []entity.Recipe
	for result.Next() {
		recipe := entity.Recipe{}
		err := result.Scan(&recipe.Id, &recipe.Title, &recipe.Ingredients, &recipe.Calories, &recipe.UserId, &recipe.CategoryId)
		helper.PanicError(err)

		recipes = append(recipes, recipe)
	}

	return recipes
}

func (repo *RecipeRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Recipe, error) {
	script := "SELECT id,title,ingredients,calories,user_id,category_id FROM recipes WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	recipe := entity.Recipe{}
	if result.Next() {
		err := result.Scan(&recipe.Id, &recipe.Title, &recipe.Ingredients, &recipe.Calories, &recipe.UserId, &recipe.CategoryId)
		helper.PanicError(err)
		return recipe, nil
	}

	return recipe, errors.New("ID not found")
}

func (repo *RecipeRepositoryImpl) Search(ctx context.Context, tx *sql.Tx, keyword string) ([]entity.Recipe, error) {
	script := "SELECT id,title,ingredients,calories,user_id,category_id FROM recipes WHERE title LIKE (?)"
	param := "%" + keyword + "%"
	result, err := tx.QueryContext(ctx, script, param)
	helper.PanicError(err)

	defer result.Close()

	var recipes []entity.Recipe
	for result.Next() {
		recipe := entity.Recipe{}
		err := result.Scan(&recipe.Id, &recipe.Title, &recipe.Ingredients, &recipe.Calories, &recipe.UserId, &recipe.CategoryId)
		helper.PanicError(err)
		recipes = append(recipes, recipe)
	}

	if len(recipes) == 0 {
		return recipes, errors.New("No Data Found")
	}

	return recipes, nil
}

func (repo *RecipeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int32) error {
	script := "DELETE FROM recipes WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, id)
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)

	if row == 0 {
		return errors.New("ID not found")
	}

	return nil

}
