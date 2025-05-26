package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Column         []string
	Id             int
	Name           string
	Email          string
	Password       string
	IsAdmin        int
	IsVerify       int
	VerifyToken    string
	ExpiredAt      time.Time
	ResetToken     sql.NullString
	ResetExpiredAt sql.NullTime
}

func (user *User) GetEntityName() string {
	return "users"

}

func (user *User) GetColumn() []string {
	return user.Column
}

func (user *User) GetId() int {
	return user.Id
}

func (user *User) SetId(id int) {
	user.Id = id
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) SetName(name string) {
	user.Name = name
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user *User) Clone() interface{} {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
