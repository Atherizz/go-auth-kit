package web

type EntityRequest interface {
	GetEntityName() string
	GetId() int
	SetId(id int)
	GetName() string
	SetName(name string)
}

type EntityResponse interface {
	GetEntityName() string
	SetName(name string)
	SetId(id int)
}
