package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	script := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, script, category.Name)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)
	return category
}

func (repo *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	script := "SELECT * FROM category"
	result, err := tx.QueryContext(ctx, script)
	helper.PanicError(err)

	defer result.Close()

	var categories []entity.Category
	for result.Next() {
		category := entity.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repo *CategoryRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (entity.Category, error) {
	script := "SELECT * FROM category WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	category := entity.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		return category, nil
	}

	return category, errors.New("ID not found")
}

func (repo *CategoryRepositoryImpl) Search(ctx context.Context, tx *sql.Tx, keyword string) ([]entity.Category, error) {
	script := "SELECT * FROM category WHERE name LIKE (?)"
	param := "%" + keyword + "%"
	result, err := tx.QueryContext(ctx, script, param)
	helper.PanicError(err)

	defer result.Close()

	var categories []entity.Category
	for result.Next() {
		category := entity.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return categories, errors.New("No Data Found")
	}

	return categories, nil
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id int, name string) (entity.Category, error) {
	script := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, name, id)
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)
	if row == 0 {
		return entity.Category{}, errors.New("no row affected")
	}

	category, _ := repo.GetById(ctx, tx, id)
	return category, nil

}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int32) error {
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
