package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	generic_models "wabustock/generics/generic-models"
)

var DB *gorm.DB

func ConnectToDB() (db *gorm.DB) {
	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return DB
}

// CreateTenantSchema creates a new schema for the tenant
func CreateTenantSchema(db *gorm.DB, tenant generic_models.Tenant) error {
	if err := db.Exec(fmt.Sprintf("CREATE SCHEMA %s", tenant.SchemaName)).Error; err != nil {
		return err
	}
	return nil
}
