package entity

type NamedEntity interface {
	GetEntityName() string
	GetId() int
	SetId(id int) 
	GetName() string
	SetName(name string) 
	
}


