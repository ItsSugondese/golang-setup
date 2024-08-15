package tenant_middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	generic_models "wabustock/generics/generic-models"
	"wabustock/pkg/common/database"
)

type tenantKey struct{}

// TenantMiddleware extracts the tenant ID from the request and adds it to the context
func TenantMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("Origin") // Or get it from URL parameters, etc.
		if tenantID == "" {
			panic("Tenant ID is required")
		}

		errVal := SetTenantSchema(database.DB, "public")
		if errVal != nil {
			panic(errVal)
		}
		tenant := generic_models.Tenant{}
		if err := db.First(&tenant, "name = ?", tenantID).Error; err != nil {
			panic("Tenant not found")
		}

		err := SetTenantSchema(database.DB, tenant.SchemaName)
		if err != nil {
			panic(err)
		}
		ctx := context.WithValue(c.Request.Context(), tenantKey{}, tenant.SchemaName)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GetTenantSchema retrieves the tenant schema from the context
func GetTenantSchema(ctx context.Context) string {
	if schema, ok := ctx.Value(tenantKey{}).(string); ok {
		return schema
	}
	return ""
}

// SetTenantSchema sets the search path for the tenant schema
func SetTenantSchema(db *gorm.DB, schemaName string) error {
	return db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error
}
