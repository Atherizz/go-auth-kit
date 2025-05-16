package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, email string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
