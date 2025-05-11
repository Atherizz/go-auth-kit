package web

type EntityRequest interface {
	GetEntityName() string
	GetId() int
	SetId(id int)
	GetName() string
	SetName(name string)
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
}

type EntityResponse interface {
	GetEntityName() string
	SetName(name string)
	SetId(id int)
	SetEmail(email string)
	SetPassword(password string)
}
