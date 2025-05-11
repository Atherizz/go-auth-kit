package web

type CategoryResponse struct {
	Entity string `json:"-"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	// Email string `json:"email"`
	// Password string `json:"-"`
}

func (model *CategoryResponse) GetEntityName() string {
	return model.Entity
}

func (model *CategoryResponse) SetId(id int) {
	model.Id = id
}

func (model *CategoryResponse) SetName(name string) {
	model.Name = name
}

func (model *CategoryResponse) SetEmail(email string) {
}

func (model *CategoryResponse) SetPassword(password string) {

}




