package models

import "gorm.io/gorm"

/**
 * Tenant represents the tenant database model.
 * Because it has a definition of gorm.Model inside of the structure, it is used for defining the Tenants table inside the database.
 */
type Tenant struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}
