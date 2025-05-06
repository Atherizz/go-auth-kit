package web

type CategoryUpdateRequest struct {
	Entity    string
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}

func(category *CategoryUpdateRequest) GetEntityName() string {
	return category.Entity
}

func(category *CategoryUpdateRequest) GetId() int {
	return category.Id
}

func (category *CategoryUpdateRequest) SetId(id int)  {
	category.Id = id
}

func(category *CategoryUpdateRequest) GetName() string {
	return category.Name
}

func(category *CategoryUpdateRequest) SetName(name string) {
	category.Name = name
}





