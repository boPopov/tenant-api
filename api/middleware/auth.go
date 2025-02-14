package middleware

import (
	"github.com/boPopov/tenant-api/api/utils"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

// Middleware to Protect API Endpoints
func JWTProtected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey: []byte(utils.JwtSecret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
	})
}
