package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}
