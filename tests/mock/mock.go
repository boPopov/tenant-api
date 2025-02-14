package mocks

import (
	"time"

	"github.com/boPopov/tenant-api/api/utils"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(utils.JwtSecret)

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

// OAuthMockGenerateToken creates a mock OAuth 2.0 token for integration testing
func OAuthMockGenerateToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"iss":      "mock-oauth-provider",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte(utils.JwtSecret)) // Same secret as the API
	return jwtToken
}
