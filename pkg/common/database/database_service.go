package database

import (
	"fmt"
	"gorm.io/gorm"
)

// SetTenantSchema sets the search path for the tenant schema
func SetTenantSchema(db *gorm.DB, schemaName string) error {
	return db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error
}
