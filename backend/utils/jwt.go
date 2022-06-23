package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

var jwtkey []byte = []byte("secret")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	return token.SignedString(jwtkey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, domain.ErrInvalidToken
		}
		return jwtkey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	return &claims, nil
}
