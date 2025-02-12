package handler

import (
	"github.com/boPopov/tenant-api/api/database"
	"github.com/boPopov/tenant-api/api/models"
	"github.com/gofiber/fiber/v2"
)

// CreateTenant - Create a new tenant
func CreateTenant(c *fiber.Ctx) error {
	tenant := new(models.Tenant)
	if err := c.BodyParser(tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&tenant)
	return c.Status(fiber.StatusCreated).JSON(tenant)
}

// GetTenant - Retrieve a tenant by ID
func GetTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	var tenant models.Tenant
	if err := database.DB.First(&tenant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}
	return c.JSON(tenant)
}

// GetAllTenants - Retrieve all tenants
func GetAllTenants(c *fiber.Ctx) error {
	var tenants []models.Tenant
	database.DB.Find(&tenants)
	return c.JSON(tenants)
}

// UpdateTenant - Update a tenant
func UpdateTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	//Check if ID is a number
	var tenant models.Tenant
	if err := database.DB.First(&tenant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	if err := c.BodyParser(&tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Save(&tenant)
	return c.JSON(tenant)
}

// DeleteTenant - Delete a tenant
func DeleteTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	//Check if id is a number
	if err := database.DB.Delete(&models.Tenant{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
