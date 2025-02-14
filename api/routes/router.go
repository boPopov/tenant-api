package routes

import (
	handler "github.com/boPopov/tenant-api/api/handlers"
	"github.com/boPopov/tenant-api/api/middleware"
	_ "github.com/boPopov/tenant-api/api/swaggerdocs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetRouths(app *fiber.App) {
	app.Static("/swagger", "./swaggerdocs")
	app.Get("/swagger/*", swagger.HandlerDefault)

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
