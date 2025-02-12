package routes

import (
	handler "github.com/boPopov/tenant-api/api/handlers"
	"github.com/boPopov/tenant-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetRouths(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/tenants", middleware.JWTProtected(), handler.CreateTenant)
	api.Get("/tenants/:id", middleware.JWTProtected(), handler.GetTenant)
	api.Get("/tenants", middleware.JWTProtected(), handler.GetAllTenants)
	api.Put("/tenants/:id", middleware.JWTProtected(), handler.UpdateTenant)
	api.Delete("/tenants/:id", middleware.JWTProtected(), handler.DeleteTenant)

	// OAuth GitHub Endpoints
	api.Get("/auth/github/login", handler.GithubLoginHandler)
	api.Get("/auth/github/callback", handler.GithubCallbackHandler)

}
