package web

type LoginResponse struct {
	Data    UserResponse `json:"data"`
	Token string `json:"token"`
}
