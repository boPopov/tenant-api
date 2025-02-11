package routes

import (
	handler "github.com/boPopov/tenant-api/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRouths(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/tenants", handler.CreateTenant)
	api.Get("/tenants/:id", handler.GetTenant)
	api.Get("/tenants", handler.GetAllTenants)
	api.Put("/tenants/:id", handler.UpdateTenant)
	api.Delete("/tenants/:id", handler.DeleteTenant)
}
