package web

type ResetPasswordRequest struct {
	NewPassword string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=NewPassword"`
}