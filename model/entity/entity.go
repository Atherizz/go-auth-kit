package entity

type NamedEntity interface {
	GetEntityName() string
	GetColumn() []string
	GetId() int
	SetId(id int)
	GetName() string
	SetName(name string)
	Clone() interface{}
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
}


