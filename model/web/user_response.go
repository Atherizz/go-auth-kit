package web

import "time"

type UserResponse struct {
	Entity      string `json:"-"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	IsVerify    int    `json:"verify_status"`
	VerifyToken string `json:"verify_token"`
	ExpiredAt   time.Time `json:"token_expired_at"`
	ResetToken string `json:"reset_token"`
	ResetExpiredAt time.Time `json:"reset_token_expired_at"`
}

func (model *UserResponse) GetEntityName() string {
	return model.Entity
}

func (model *UserResponse) SetId(id int) {
	model.Id = id
}

func (model *UserResponse) SetName(name string) {
	model.Name = name
}

func (model *UserResponse) SetEmail(email string) {
	model.Email = email
}

func (model *UserResponse) SetPassword(password string) {
	model.Password = password
}
