package helper

import (
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
)

func ToEntityResponse[S entity.NamedEntity, R web.EntityResponse](model S, constructor func() R) R {
	response := constructor()
	response.SetId(model.GetId())
	response.SetName(model.GetName())
	response.SetEmail(model.GetEmail())

	return response
}

func ToUserResponse(user entity.User) web.UserResponse {
	return web.UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		IsVerify:    user.IsVerify,
		VerifyToken: user.VerifyToken,
		ExpiredAt:   user.ExpiredAt,
		ResetToken:     user.ResetToken.String,
		ResetExpiredAt: user.ResetExpiredAt.Time,
	}
}

func ToVerifyTokenResponse(user entity.User) web.VerifyTokenResponse {
	return web.VerifyTokenResponse{
		Email:       user.Email,
		IsVerify:    user.IsVerify,
		VerifyToken: user.VerifyToken,
		ExpiredAt:   user.ExpiredAt,
	}
}

func ToResetTokenResponse(user entity.User) web.ResetTokenResponse {
	return web.ResetTokenResponse{
		Email:          user.Email,
		ResetToken:     user.ResetToken.String,
		ResetExpiredAt: user.ResetExpiredAt.Time,
	}
}

func ToRecipeResponse(recipe entity.Recipe) web.RecipeResponse {
	return web.RecipeResponse{
		Id:          recipe.Id,
		Title:       recipe.Title,
		Ingredients: recipe.Ingredients,
		Calories:    float64(recipe.Calories),
		UserId:      recipe.UserId,
		CategoryId:  recipe.CategoryId,
	}
}
