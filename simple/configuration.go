package simple

type Configuration struct {
	Name string
}



type Apllication struct {
	*Configuration
}

func NewApplication() *Apllication {
	return &Apllication{
		Configuration: &Configuration{
			Name: "Golang",
		},
	}
}