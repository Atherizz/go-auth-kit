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
	script := "INSERT INTO categories(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, script, model.GetName())
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	model.SetId(int(id))
	return model
}

func (repo *RepositoryImpl[T]) GetAll(ctx context.Context, tx *sql.Tx, model T) []T {
	script := "SELECT * FROM categories"
	result, err := tx.QueryContext(ctx, script)
	helper.PanicError(err)

	defer result.Close()

	var categories []T
	for result.Next() {
		category := model.Clone().(T)
		var id int
		var name string

		err := result.Scan(&id, &name)
		helper.PanicError(err)

		category.SetId(id)
		category.SetName(name)
		categories = append(categories, category)
	}

	return categories
}

func (repo *RepositoryImpl[T]) GetById(ctx context.Context, tx *sql.Tx, id int, model T) (T, error) {
	script := "SELECT * FROM categories WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	if result.Next() {
		category := model.Clone().(T)
		var id int
		var name string

		err := result.Scan(&id, &name)
		helper.PanicError(err)

		category.SetId(id)
		category.SetName(name)
		return category, nil
	}

	return model, errors.New("ID not found")
}

func (repo *RepositoryImpl[T]) Search(ctx context.Context, tx *sql.Tx, keyword string, model T) ([]T, error) {
	script := "SELECT * FROM categories WHERE name LIKE (?)"
	param := "%" + keyword + "%"
	result, err := tx.QueryContext(ctx, script, param)
	helper.PanicError(err)

	defer result.Close()

	var categories []T
	for result.Next() {
		category := model.Clone().(T)
		var id int
		var name string

		err := result.Scan(&id, &name)
		helper.PanicError(err)

		category.SetId(id)
		category.SetName(name)
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return categories, errors.New("No Data Found")
	}

	return categories, nil
}

func (repo *RepositoryImpl[T]) Update(ctx context.Context, tx *sql.Tx, model T) (T, error) {
	script := "UPDATE categories SET name = ? WHERE id = ?"
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
	script := "DELETE FROM categories WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, id)
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)

	if row == 0 {
		return errors.New("ID not found")
	}

	return nil

}
