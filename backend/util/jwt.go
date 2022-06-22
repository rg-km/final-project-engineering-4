package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int64, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte("secret"))
}
