// Package swaggerdocs Code generated by swaggo/swag. DO NOT EDIT
package swaggerdocs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.example.com/support",
            "email": "bojpopov@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/github/callback": {
            "get": {
                "description": "Handles the OAuth callback from GitHub",
                "tags": [
                    "Authentication"
                ],
                "summary": "GitHub OAuth Callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/github/login": {
            "get": {
                "description": "Redirects user to GitHub for OAuth authentication",
                "tags": [
                    "Authentication"
                ],
                "summary": "GitHub Login",
                "responses": {
                    "302": {
                        "description": "Found"
                    }
                }
            }
        },
        "/tenants": {
            "get": {
                "description": "Fetches all tenants from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenants"
                ],
                "summary": "Get all tenants",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new tenant and stores it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenants"
                ],
                "summary": "Create a new tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Tenant Request",
                        "name": "tenant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TenantRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/tenants/{id}": {
            "get": {
                "description": "Retrieves a single tenant using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenants"
                ],
                "summary": "Get a tenant by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a tenant’s details using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenants"
                ],
                "summary": "Update a tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Tenant Request",
                        "name": "tenant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TenantRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a tenant using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenants"
                ],
                "summary": "Delete a tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.TenantRequest": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean",
                    "example": true
                },
                "email": {
                    "type": "string",
                    "example": "johndoe@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tenant API",
	Description:      "This is a sample API for managing tenants.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
