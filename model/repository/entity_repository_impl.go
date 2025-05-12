package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type EntityRepositoryImpl[T entity.NamedEntity] struct {
}

func NewRepository[T entity.NamedEntity]() *EntityRepositoryImpl[T] {
	return &EntityRepositoryImpl[T]{}
}

func (repo *EntityRepositoryImpl[T]) Create(ctx context.Context, tx *sql.Tx, model T) T {
	entity := model.GetEntityName()
	getColumn := model.GetColumn()
	prepare := "("

	
	for i := 0; i < len(getColumn); i++ {
		if i != len(getColumn)-1 {
			prepare += "?,"
		} else {
			prepare += "?)"
		}
	}

	column := "(" + strings.Join(getColumn, ",") + ")"
	script := "INSERT INTO " + entity + " " + column + "VALUES " + prepare + ""
	
	switch entity {
	case "categories":
		result, err := tx.ExecContext(ctx, script, model.GetName())
		helper.PanicError(err)
		id, err := result.LastInsertId()
		helper.PanicError(err)
		model.SetId(int(id))	
	case "users" :
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.GetPassword()), bcrypt.DefaultCost)
		helper.PanicError(err)
		hashedString := string(hashedPassword)

		result, err := tx.ExecContext(ctx, script, model.GetName(), model.GetEmail(), hashedString)
		helper.PanicError(err)
		id, err := result.LastInsertId()
		helper.PanicError(err)
		model.SetId(int(id))	
	} 
	
	return model
}

func (repo *EntityRepositoryImpl[T]) GetAll(ctx context.Context, tx *sql.Tx, model T) []T {
	entity := model.GetEntityName()

	script := "SELECT * FROM " + entity + ""
	result, err := tx.QueryContext(ctx, script)
	helper.PanicError(err)

	defer result.Close()

	var entities []T
	for result.Next() {
		dataEntity := model.Clone().(T)
	switch entity {
	case "categories":
			var id int
			var name string
	
			err := result.Scan(&id, &name)
			helper.PanicError(err)
			dataEntity.SetId(id)
			dataEntity.SetName(name)
			entities = append(entities, dataEntity)

	case "users" :
			var id int
			var name string
			var email string
			var password string 

			err := result.Scan(&id, &name, &email, &password)
			helper.PanicError(err)
			dataEntity.SetId(id)
			dataEntity.SetName(name)
			dataEntity.SetEmail(email)
			dataEntity.SetPassword(password)
			entities = append(entities, dataEntity)
		}	
	} 
	
	return entities
}

func (repo *EntityRepositoryImpl[T]) GetById(ctx context.Context, tx *sql.Tx, id int, model T) (T, error) {
	entity := model.GetEntityName()

	script := "SELECT * FROM " + entity + " WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	if result.Next() {
		dataEntity := model.Clone().(T)
		switch entity {
		case "categories":
				var id int
				var name string
		
				err := result.Scan(&id, &name)
				helper.PanicError(err)
				dataEntity.SetId(id)
				dataEntity.SetName(name)
				return dataEntity, nil
	
		case "users" :
				var id int
				var name string
				var email string
				var password string 
	
				err := result.Scan(&id, &name, &email, &password)
				helper.PanicError(err)
				dataEntity.SetId(id)
				dataEntity.SetName(name)
				dataEntity.SetEmail(email)
				dataEntity.SetPassword(password)
				return dataEntity, nil
			}	
	}

	return model, errors.New("ID not found")
}


func (repo *EntityRepositoryImpl[T]) Search(ctx context.Context, tx *sql.Tx, keyword string, model T) ([]T, error) {
	entity := model.GetEntityName()

	script := "SELECT * FROM  " + entity + " WHERE name LIKE (?)"
	param := "%" + keyword + "%"
	result, err := tx.QueryContext(ctx, script, param)
	helper.PanicError(err)

	defer result.Close()

	var entities []T
	for result.Next() {
		dataEntity := model.Clone().(T)
	switch entity {
	case "categories":
			var id int
			var name string
	
			err := result.Scan(&id, &name)
			helper.PanicError(err)
			dataEntity.SetId(id)
			dataEntity.SetName(name)
			entities = append(entities, dataEntity)
	case "users" :
			var id int
			var name string
			var email string
			var password string 

			err := result.Scan(&id, &name, &email, &password)
			helper.PanicError(err)
			dataEntity.SetId(id)
			dataEntity.SetName(name)
			dataEntity.SetEmail(email)
			dataEntity.SetPassword(password)
			entities = append(entities, dataEntity)
		}	
	} 
	
	if len(entities) == 0 {
		return entities, errors.New("No Data Found")
	}

	return entities, nil
}

func (repo *EntityRepositoryImpl[T]) Update(ctx context.Context, tx *sql.Tx, model T) (T, error) {
	entity := model.GetEntityName()

	script := "UPDATE " + entity + " SET name = ? WHERE id = ?"
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

func (repo *EntityRepositoryImpl[T]) Delete(ctx context.Context, tx *sql.Tx, id int32, model T) error {
	entity := model.GetEntityName()

	script := "DELETE FROM " + entity + " WHERE id = ?"
	result, err := tx.ExecContext(ctx, script, id)
	helper.PanicError(err)

	row, err := result.RowsAffected()
	helper.PanicError(err)

	if row == 0 {
		return errors.New("ID not found")
	}

	return nil
}
