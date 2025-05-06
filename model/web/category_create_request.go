package web

type CategoryCreateRequest struct {
	Entity    string
	Name string `json:"name" validate:"required,min=1,max=200"`
}

func(category *CategoryCreateRequest) GetEntityName() string {
	return category.Entity
}

func(category *CategoryCreateRequest) GetId() int {
	return -1
}

func (category *CategoryCreateRequest) SetId(id int)  {
	
}

func(category *CategoryCreateRequest) GetName() string {
	return category.Name
}

func(category *CategoryCreateRequest) SetName(name string) {
	category.Name = name
}
