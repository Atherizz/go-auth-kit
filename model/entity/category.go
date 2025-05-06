package entity

type Category struct {
	Entity string
	Id     int
	Name   string
}

func (category *Category) GetEntityName() string {
	return category.Entity
}

func (category *Category) GetId() int {
	return category.Id
}

func (category *Category) SetId(id int) {
	category.Id = id
}

func (category *Category) GetName() string {
	return category.Name
}

func (category *Category) SetName(name string) {
	category.Name = name
}

func (category *Category) Clone() interface{} {
	return &Category{
		Id:   category.Id,
		Name: category.Name,
	}
}
