package tenant

import (
	"fmt"
	"gorm.io/gorm"
	tenant_constants "wabustock/constants/tenant-constants"
	"wabustock/pkg/common/database"
)

// CreateTenantSchema creates a new schema for the tenant
func CreateTenantSchemaRepo(db *gorm.DB, tenant Tenant) error {
	if err := db.Exec(fmt.Sprintf("CREATE SCHEMA %s", tenant.Name)).Error; err != nil {
		return err
	}
	return nil
}

func SaveTenantDetails(db *gorm.DB, tenant Tenant) error {
	if err := database.SetTenantSchema(db, tenant_constants.PublicTenant); err != nil {
		return err
	}
	result := database.DB.Create(&tenant)
	return result.Error
}
