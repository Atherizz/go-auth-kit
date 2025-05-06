package web

type CategoryResponse struct {
	Entity string
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

func (category *CategoryResponse) GetEntityName() string {
	return category.Entity
}

func (category *CategoryResponse) SetId(id int) {
	category.Id = id
}

func (category *CategoryResponse) SetName(name string) {
	category.Name = name
}
