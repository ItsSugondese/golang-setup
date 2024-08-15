package generic_models

import (
	"gorm.io/gorm"
)

// Tenant represents a tenant in your system
type Tenant struct {
	gorm.Model
	Name       string `gorm:"uniqueIndex"`
	SchemaName string
}
