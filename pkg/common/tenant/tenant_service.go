package tenant

import (
	"gorm.io/gorm"
	tenant_constants "wabustock/constants/tenant-constants"
	generic_models "wabustock/generics/generic-models"
	"wabustock/internal/role"
	temporary_attachments "wabustock/internal/temporary-attachments"
	"wabustock/internal/user"
	"wabustock/pkg/common/database"
)

// CreateTenantSchema creates a new schema for the tenant
func CreateTenantSchemaService(dto TenantRequest) string {

	tenant := Tenant{
		Name: dto.Name,
	}

	err := CreateTenantSchemaRepo(database.DB, tenant)
	if err != nil {
		panic(err)
	}

	migrateTenantErr := MigrateTenantTables(database.DB, tenant)
	if migrateTenantErr != nil {
		panic(migrateTenantErr)
	}

	setTenantErr := database.SetTenantSchema(database.DB, tenant_constants.PublicTenant)
	if setTenantErr != nil {
		panic(setTenantErr)
	}

	saveTenantError := SaveTenantDetails(database.DB, tenant)
	if saveTenantError != nil {
		panic(saveTenantError)
	}
	return tenant.Name
}

// MigrateTenantTables migrates the necessary tables within the tenant schema
func MigrateTenantTables(db *gorm.DB, tenant Tenant) error {
	if err := database.SetTenantSchema(db, tenant.Name); err != nil {
		return err
	}

	//Automigrate tenant-specific tables
	if err := db.AutoMigrate(&temporary_attachments.TemporaryAttachments{}); err != nil { // Add more models as needed
		return err
	}

	if err := db.AutoMigrate(&user.BaseUser{}, &generic_models.AuditModel{}); err != nil { // Add more models as needed
		return err
	}
	if err := db.AutoMigrate(&role.Role{}); err != nil { // Add more models as needed
		return err
	}

	return nil
}

func MigrateTenantPublicTable(db *gorm.DB) error {
	publicTenantExists, queryError := database.SchemaExistsRepo(db, tenant_constants.PublicTenant)
	if queryError != nil {
		panic(queryError)
	}

	if !publicTenantExists {

		publicTenantCreateError := CreateTenantSchemaRepo(database.DB, Tenant{
			Name: tenant_constants.PublicTenant,
		})

		if publicTenantCreateError != nil {
			panic(publicTenantCreateError)
		}

	}
	if err := database.SetTenantSchema(db, tenant_constants.PublicTenant); err != nil {
		return err
	}

	//Automigrate tenant-specific tables
	if err := db.AutoMigrate(&Tenant{}); err != nil { // Add more models as needed
		return err
	}
	return nil
}

func MigrateAll() error {
	//Automigrate tenant-specific tables
	if err := database.DB.AutoMigrate(&temporary_attachments.TemporaryAttachments{}, &user.BaseUser{}, &generic_models.AuditModel{}, &role.Role{}); err != nil { // Add more models as needed
		return err
	}
	return nil
}
