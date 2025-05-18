package entity

type Recipe struct {
	Id          int
	Title       string
	Ingredients string
	Calories    float64
	UserId      int
	CategoryId 	int
}
