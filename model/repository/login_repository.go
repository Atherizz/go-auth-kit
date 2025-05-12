package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
)

type LoginRepository interface {
	GetByName(ctx context.Context, tx *sql.Tx, name string) (entity.User,error)
}