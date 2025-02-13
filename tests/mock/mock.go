package mocks

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("supersecretkey")

// MockGenerateJWT creates a simple mock JWT token for testing
func MockGenerateJWT(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 1-hour expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString(jwtSecret)
	return jwtToken
}
