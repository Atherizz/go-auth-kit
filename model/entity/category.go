package entity

type Category struct {
	Column []string
	Id     int
	Name   string
}

func (category *Category) GetEntityName() string {
	return "categories"
}

func (category *Category) GetColumn() []string {
	return category.Column
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

func (category *Category) GetEmail() string {
	return ""
}

func (category *Category) SetEmail(email string) {
	
}

func (category *Category) GetPassword() string {
	return ""
}

func (category *Category) SetPassword(password string) {
	
}

func (category *Category) Clone() interface{} {
	return &Category{
		Id:   category.Id,
		Name: category.Name,
	}
}
