package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/entity"
)

type AuthRepository interface {
	GetByColumn(ctx context.Context, tx *sql.Tx, data string, column string) (entity.User,error)
	GetById(ctx context.Context, tx *sql.Tx, id int) (entity.User,error)
	Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User)
	SetVerified(ctx context.Context, tx *sql.Tx, token string) (entity.User,error)
}