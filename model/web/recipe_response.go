package web

type RecipeResponse struct {
	Id int 				`json:"id" validate:"required,number"`
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Ingredients string `json:"ingredients" validate:"required,min=1,max=200"`
	Calories    float64    `json:"calories" validate:"required,number"`
	UserId      int    `json:"user_id" validate:"required,number"`
	CategoryId  int    `json:"category_id" validate:"required,number"`
}