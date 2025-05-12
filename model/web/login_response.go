package web

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    UserResponse `json:"data"`
	Token string `json:"token"`
}
