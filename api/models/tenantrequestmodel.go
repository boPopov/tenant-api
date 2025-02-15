package models

/**
 * TenantRequest represents the request body for creating a tenant.
 * This Tenant is main purpose is to extract the body from the API calls.
 * Additionally it is a solution to display the Example body in the Swagger Documentation
 */
type TenantRequest struct {
	Name   string `json:"name" example:"John Doe"`
	Email  string `json:"email" example:"johndoe@example.com"`
	Active bool   `json:"active" example:"true"`
}
