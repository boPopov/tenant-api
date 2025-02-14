package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}
