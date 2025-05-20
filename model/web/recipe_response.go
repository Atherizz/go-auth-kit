package web

type RecipeResponse struct {
	Id int 				`json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Calories    float64    `json:"calories"`
	UserId      int    `json:"user_id"`
	CategoryId  int    `json:"category_id"`
}