package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// OAuth Config
var oauthConfig = &oauth2.Config{
	ClientID:     "Ov23liqCWXLsBMtqEUaM",                     // Read from env
	ClientSecret: "a14ea7a80a8720de2683e1f1e2eedc3f8459a40a", // Read from env
	RedirectURL:  "http://localhost:3000/api/auth/github/callback",
	Scopes:       []string{"user"},
	Endpoint:     github.Endpoint,
}

var jwtSecret = []byte("supersecretkey") // Change this in production!

// GitHub Login Handler - Redirects to GitHub OAuth
func GithubLoginHandler(c *fiber.Ctx) error {
	url := oauthConfig.AuthCodeURL("randomStringForCSRF", oauth2.AccessTypeOffline)
	return c.JSON(fiber.Map{"auth_url": url})
}

// GitHub Callback Handler - Exchanges code for access token
func GithubCallbackHandler(c *fiber.Ctx) error {
	// Get code from GitHub callback
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing code parameter"})
	}

	// Exchange the authorization code for an access token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// Fetch user details from GitHub
	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	// Generate JWT Token
	jwtToken, err := generateJWT(user["login"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Return JWT Token
	return c.JSON(fiber.Map{
		"access_token": jwtToken,
		"user": fiber.Map{
			"username": user["login"],
			"email":    user["email"],
			"profile":  user["html_url"],
		},
	})
}

// Generate JWT Token
func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
