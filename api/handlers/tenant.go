package handler

import (
	"github.com/boPopov/tenant-api/api/database"
	"github.com/boPopov/tenant-api/api/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create a new tenant
// @Description Creates a new tenant and stores it in the database
// @Tags Tenants
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Param tenant body models.TenantRequest true "Tenant Request"
// @Produce json
// @Success 201 {object} object
// @Router /tenants [post]
func CreateTenant(c *fiber.Ctx) error {
	var request models.TenantRequest

	// Parse and validate request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Create a Tenant record from the request
	tenant := models.Tenant{
		Name:   request.Name,
		Email:  request.Email,
		Active: request.Active,
	}

	if tenant.Name == "" && tenant.Email == "" { //Checking if email and/or name are not set.
		return c.Status(fiber.StatusBadRequest).JSON("You must provide a value for Name and/or Email!")
	}

	database.DB.Create(&tenant) //Creating a new record in database for the new entered tenant

	return c.Status(fiber.StatusCreated).JSON(tenant) //Returning status created.
}

// @Summary Get a tenant by ID
// @Description Retrieves a single tenant using its ID
// @Tags Tenants
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Param id path string true "Tenant ID"
// @Success 200 {object} object
// @Router /tenants/{id} [get]
func GetTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	var tenant models.Tenant
	if err := database.DB.First(&tenant, id).Error; err != nil { //If tenant is not found in Database, we return  404.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}
	return c.JSON(tenant)
}

// @Summary Get all tenants
// @Description Fetches all tenants from the database
// @Tags Tenants
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {array} object
// @Router /tenants [get]
func GetAllTenants(c *fiber.Ctx) error {
	var tenants []models.Tenant
	database.DB.Find(&tenants)
	return c.JSON(tenants)
}

// @Summary Update a tenant
// @Description Updates a tenant’s details using its ID
// @Tags Tenants
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Param tenant body models.TenantRequest true "Tenant Request"
// @Produce json
// @Param id path string true "Tenant ID"
// @Success 200 {object} object
// @Router /tenants/{id} [put]
func UpdateTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	//Check if ID is a number
	var tenant models.Tenant
	if err := database.DB.First(&tenant, id).Error; err != nil { //Finding tenant.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	if err := c.BodyParser(&tenant); err != nil { //Checking the parsing
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Save(&tenant) //Saving in DB
	return c.JSON(tenant)
}

// @Summary Delete a tenant
// @Description Deletes a tenant using its ID
// @Tags Tenants
// @Accept json
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Param id path string true "Tenant ID"
// @Success 204
// @Router /tenants/{id} [delete]
func DeleteTenant(c *fiber.Ctx) error {
	id := c.Params("id")

	var tenant models.Tenant
	if err := database.DB.First(&tenant, id).Error; err != nil { //Finding tenant.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	if err := database.DB.Delete(&tenant, id).Error; err != nil { //Deleting tenant
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
