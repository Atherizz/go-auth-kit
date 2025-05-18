package web

type RecipeRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Ingredients string `json:"ingredients" validate:"required,min=1,max=200"`
	Calories    int    `json:"calories" validate:"required,number"`
	UserId      int    `json:"user_id" validate:"required,number"`
	CategoryId  int    `json:"category_id" validate:"required,number"`
}
