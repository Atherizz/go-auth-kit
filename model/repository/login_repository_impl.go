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

func NewLoginRepository () *LoginRepositoryImpl {
	return &LoginRepositoryImpl{}
}

func (repo *LoginRepositoryImpl) GetByName(ctx context.Context, tx *sql.Tx, name string) (entity.User, error) {
	script := "SELECT * FROM users WHERE name = (?)"
	result, err := tx.QueryContext(ctx, script, name)
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
