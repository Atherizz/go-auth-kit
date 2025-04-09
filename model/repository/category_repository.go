package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Category
	GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Category, error)
	Update(ctx context.Context, tx *sql.Tx, id int, name string) (entity.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, id int32) (error)
	Search(ctx context.Context, tx *sql.Tx, keyword string) ([]entity.Category, error) 
}
