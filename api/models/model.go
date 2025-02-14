package models

import "gorm.io/gorm"

// Tenant represents the tenant database model
type Tenant struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// TenantRequest represents the request body for creating a tenant
type TenantRequest struct {
	Name   string `json:"name" example:"John Doe"`
	Email  string `json:"email" example:"johndoe@example.com"`
	Active bool   `json:"active" example:"true"`
}
