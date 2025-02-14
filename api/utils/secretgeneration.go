package utils

import (
	"crypto/rand"
	"encoding/hex"
)

var JwtSecret string

// GenerateSecureJWTSecret creates a 32-byte secure random key
func GenerateSecureJWTSecret() {
	bytes := make([]byte, 32) // 256-bit key
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to generate secure random bytes: " + err.Error())
	}
	JwtSecret = hex.EncodeToString(bytes)
}
