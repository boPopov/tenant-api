package mocks

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("544c142b11d962494bc6d7ecffcd53c94862bb849f0c245ae8bb9715b10a03c6")

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
	jwtToken, _ := token.SignedString([]byte("544c142b11d962494bc6d7ecffcd53c94862bb849f0c245ae8bb9715b10a03c6")) // Same secret as the API
	return jwtToken
}
