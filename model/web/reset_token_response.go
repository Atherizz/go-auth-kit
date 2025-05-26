package web

import "time"

type ResetTokenResponse struct {
	Email          string    `json:"email"`
	ResetToken string `json:"reset_token"`
	ResetExpiredAt time.Time `json:"reset_token_expired_at"`
}