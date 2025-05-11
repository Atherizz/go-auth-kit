package web

type CategoryRequest struct {
	Entity string `json:"-"`
	Id     int    `json:"id"`
	Name   string `json:"name" validate:"required,min=1,max=200"`
}

func (category *CategoryRequest) GetEntityName() string {
	return category.Entity
}

func (category *CategoryRequest) GetId() int {
	return category.Id
}

func (category *CategoryRequest) SetId(id int) {
	category.Id = id
}

func (category *CategoryRequest) GetName() string {
	return category.Name
}

func (category *CategoryRequest) SetName(name string) {
	category.Name = name
}

func (category *CategoryRequest) GetEmail() string {
	return ""
}

func (category *CategoryRequest) SetEmail(email string) {

}

func (category *CategoryRequest) GetPassword() string {
	return ""
}

func (category *CategoryRequest) SetPassword(password string) {

}




