package utils

import (
	"Taskie/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
