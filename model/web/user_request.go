package web

type UserRequest struct {
	Entity string `json:"-"`
	Id     int    `json:"id"`
	Name   string `json:"name" validate:"required,min=1,max=200"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (entity *UserRequest) GetEntityName() string {
	return entity.Entity
}

func (entity *UserRequest) GetId() int {
	return entity.Id
}

func (entity *UserRequest) SetId(id int) {
	entity.Id = id
}

func (entity *UserRequest) GetName() string {
	return entity.Name
}

func (entity *UserRequest) SetName(name string) {
	entity.Name = name
}

func (entity *UserRequest) GetEmail() string {
	return entity.Email
}

func (entity *UserRequest) SetEmail(email string) {
	entity.Email = email
}

func (entity *UserRequest) GetPassword() string {
	return entity.Password
}

func (entity *UserRequest) SetPassword(password string) {
	entity.Password = password
}




