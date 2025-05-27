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
	ResendVerifyToken(ctx context.Context, tx *sql.Tx, email string) (entity.User,error)
	ForgotPassword(ctx context.Context, tx *sql.Tx, email string) (entity.User,error)
 	ResetPassword(ctx context.Context, tx *sql.Tx, newPassword string, token string) error
	ChangePassword(ctx context.Context, tx *sql.Tx, newPassword string, id int) error
}