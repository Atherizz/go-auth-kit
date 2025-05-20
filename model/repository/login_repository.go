package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
)

type LoginRepository interface {
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User,error)
	GetById(ctx context.Context, tx *sql.Tx, id int) (entity.User,error)
}