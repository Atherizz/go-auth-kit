package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
)

type LoginRepositoryImpl struct {
}

func NewLoginRepository() *LoginRepositoryImpl {
	return &LoginRepositoryImpl{}
}

func (repo *LoginRepositoryImpl) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	script := "SELECT id,name,email,password_hash FROM users WHERE email = (?)"
	result, err := tx.QueryContext(ctx, script, email)
	helper.PanicError(err)

	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicError(err)
		return user, nil
	}

	return user, errors.New("ID not found")
}

func (repo *LoginRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	script := "SELECT id,name,email,password_hash,is_admin FROM users WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.IsAdmin)
		helper.PanicError(err)
		return user, nil
	}

	return user, errors.New("ID not found")
}
