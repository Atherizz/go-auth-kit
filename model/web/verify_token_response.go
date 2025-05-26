package web

import "time"

type VerifyTokenResponse struct {
	Email          string    `json:"email"`
	IsVerify       int       `json:"verify_status"`
	VerifyToken    string    `json:"verify_token"`
	ExpiredAt      time.Time `json:"token_expired_at"`
}