package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

// Middleware to Protect API Endpoints
func JWTProtected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey: []byte("supersecretkey"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
	})
}
