package web

type UserResponse struct {
	Entity string `json:"-"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
	RegisterToken string `json:"register_token"`
	VerifyToken string `json:"verify_token"`
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




