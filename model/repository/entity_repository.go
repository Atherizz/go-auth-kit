package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
	// "golang-restful-api/model/entity"
)

type Repository[T entity.NamedEntity] interface {
	Create(ctx context.Context, tx *sql.Tx, model T) T
	GetAll(ctx context.Context, tx *sql.Tx, model T) []T
	GetById(ctx context.Context, tx *sql.Tx, id int, model T) (T, error)
	Update(ctx context.Context, tx *sql.Tx,  model T) (T, error)
	Delete(ctx context.Context, tx *sql.Tx, id int32, model T) error
	Search(ctx context.Context, tx *sql.Tx, keyword string, model T) ([]T, error)

}
