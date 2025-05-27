package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/helper"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (repo *AuthRepositoryImpl) GetByColumn(ctx context.Context, tx *sql.Tx, data string, column string) (entity.User, error) {
script := "SELECT id,name,email,password_hash,is_admin,is_verified,verify_token,token_expired_at,reset_token,reset_token_expired_at FROM users WHERE " + column + " = ?"
	result, err := tx.QueryContext(ctx, script, data)
	helper.PanicError(err)

	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.IsAdmin, &user.IsVerify, &user.VerifyToken, &user.ExpiredAt, &user.ResetToken, &user.ResetExpiredAt)
		if err != nil {
			return user, err
		}
		return user, nil
	}

	return user, errors.New("ID not found")
}

func (repo *AuthRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	script := "SELECT id,name,email,password_hash,is_admin,is_verified,verify_token,token_expired_at,reset_token,reset_token_expired_at FROM users WHERE id = (?)"
	result, err := tx.QueryContext(ctx, script, id)
	helper.PanicError(err)

	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.IsAdmin, &user.IsVerify, &user.VerifyToken, &user.ExpiredAt, &user.ResetToken, &user.ResetExpiredAt)
		helper.PanicError(err)
		return user, nil
	}

	return user, errors.New("ID not found")
}

func (repo *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	script := "INSERT INTO users(name,email,password_hash,verify_token,token_expired_at) VALUES (?,?,?,?,?)"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helper.PanicError(err)
	hashedString := string(hashedPassword)
	// expiredAt := time.Now().Add(15 * time.Minute)

	result, err := tx.ExecContext(ctx, script, user.Name, user.Email, hashedString, user.VerifyToken, user.ExpiredAt)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	user.Id = int(id)
	return user
}

func (repo *AuthRepositoryImpl) SetVerified(ctx context.Context, tx *sql.Tx, token string) (entity.User, error) {
	script := "UPDATE users SET is_verified = 1 WHERE verify_token = (?)"
	result, err := tx.ExecContext(ctx, script, token)
	if err != nil {
		log.Printf("Error executing update: %v", err)
		return entity.User{}, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}

	if row == 0 {
		return entity.User{}, errors.New("no row affected")
	}

	res, err := repo.GetByColumn(ctx, tx, token, "verify_token")
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return entity.User{}, err
	}
	log.Printf("User retrieved after verification: %+v", res)
	return res, nil
}

func (repo *AuthRepositoryImpl) ResendVerifyToken(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	script := "UPDATE users SET verify_token = (?), token_expired_at = (?) WHERE email = (?)"

	token := uuid.NewString()
	expiredAt := time.Now().Add(15 * time.Minute)

	result, err := tx.ExecContext(ctx, script, token, expiredAt, email)
	if err != nil {
		return entity.User{}, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}

	if row == 0 {
		return entity.User{}, errors.New("no row affected")
	}

	res, err := repo.GetByColumn(ctx, tx, email, "email")
	if err != nil {
		return entity.User{}, err
	}

	return res, nil

}

func (repo *AuthRepositoryImpl) ForgotPassword(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	script := "UPDATE users SET reset_token = (?), reset_token_expired_at = (?) WHERE email = (?)"

	token := uuid.NewString()
	expiredAt := time.Now().Add(15 * time.Minute)

	result, err := tx.ExecContext(ctx, script, token, expiredAt, email)
	if err != nil {
		return entity.User{}, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}

	if row == 0 {
		return entity.User{}, errors.New("no row affected")
	}

	res, err := repo.GetByColumn(ctx, tx, email, "email")
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}

func (repo *AuthRepositoryImpl) ResetPassword(ctx context.Context, tx *sql.Tx, newPassword string, token string) error {
	script := "UPDATE users SET password_hash = (?) WHERE reset_token = (?)"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	helper.PanicError(err)
	hashedString := string(hashedPassword)

	result, err := tx.ExecContext(ctx, script,hashedString, token)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return errors.New("no row affected")
	}

	return nil
}

func (repo *AuthRepositoryImpl) ChangePassword(ctx context.Context, tx *sql.Tx, newPassword string, id int) error {
	script := "UPDATE users SET password_hash = (?) WHERE id = (?)"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	helper.PanicError(err)
	hashedString := string(hashedPassword)

	result, err := tx.ExecContext(ctx, script,hashedString, id)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("no row affected")
	}
	return nil
}
