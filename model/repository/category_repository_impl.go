package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
)

type RepositoryImpl[T entity.NamedEntity] struct {
}

func NewRepository[T entity.NamedEntity]() *RepositoryImpl[T] {
	return &RepositoryImpl[T]{}
}

func (repo *RepositoryImpl[T]) Create(ctx context.Context, tx *sql.Tx, model T) T {
	script := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, script, model.GetName())
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	model.SetId(int(id))
	return model
}

func (repo *RepositoryImpl[T]) GetAll(ctx context.Context, tx *sql.Tx, model T) []T {
	script := "SELECT * FROM category"
	result, err := tx.QueryContext(ctx, script)
	helper.PanicError(err)

	defer result.Close()

	var categories []T
	for result.Next() {
		category := model
		err := result.Scan(category.GetId(), category.GetName)
		helper.PanicError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repo *RepositoryImpl[T]) GetById(ctx context.Context, tx *sql.Tx, id int, model T) (T, error) {
	script := "SELECT * FROM category WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	category := model
	if result.Next() {
		err := result.Scan(category.GetId(), category.GetName())
		helper.PanicError(err)
		return category, nil
	}

	return category, errors.New("ID not found")
}

func (repo *RepositoryImpl[T]) Search(ctx context.Context, tx *sql.Tx, keyword string, model T) ([]T, error) {
	script := "SELECT * FROM category WHERE name LIKE (?)"
	param := "%" + keyword + "%"
	result, err := tx.QueryContext(ctx, script, param)
	helper.PanicError(err)

	defer result.Close()

	var categories []T
	for result.Next() {
		category := model
		err := result.Scan(category.GetId(), category.GetName())
		helper.PanicError(err)
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return categories, errors.New("No Data Found")
	}

	return categories, nil
}

func (repo *RepositoryImpl[T]) Update(ctx context.Context, tx *sql.Tx, model T) (T, error) {
	script := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, model.GetName(), model.GetId())
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)
	if row == 0 {
		return model, errors.New("no row affected")
	}

	res, _ := repo.GetById(ctx, tx, model.GetId(), model)
	return res, nil

}

func (repo *RepositoryImpl[T]) Delete(ctx context.Context, tx *sql.Tx, id int32) error {
	script := "DELETE FROM category WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, id)
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)

	if row == 0 {
		return errors.New("ID not found")
	}

	return nil

}
