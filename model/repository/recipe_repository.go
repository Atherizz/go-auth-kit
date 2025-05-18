package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
)

type RecipeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, recipe entity.Recipe) entity.Recipe
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Recipe
	GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Recipe, error)
	Delete(ctx context.Context, tx *sql.Tx, id int32) error
	Search(ctx context.Context, tx *sql.Tx, keyword string) ([]entity.Recipe, error)
}