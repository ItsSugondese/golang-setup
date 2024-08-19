package tenant_middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	tenant_constants "wabustock/constants/tenant-constants"
	"wabustock/pkg/common/database"
	"wabustock/pkg/common/tenant"
)

type tenantKey struct{}

// TenantMiddleware extracts the tenant ID from the request and adds it to the context
func TenantMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		errVal := database.SetTenantSchema(database.DB, tenant_constants.PublicTenant)
		if errVal != nil {
			panic(errVal)
		}

		//path := c.FullPath()
		path := c.Request.URL.Path
		if path == "/tenant" || len(path) >= 8 && path[:8] == "/tenant/" {
			c.Next() // Skip middleware logic
			return
		}
		tenantID := c.GetHeader("Origin") // Or get it from URL parameters, etc.
		if tenantID == "" {
			panic("Tenant ID is required")
		}

		tenant := tenant.Tenant{}
		if err := db.First(&tenant, "name = ?", tenantID).Error; err != nil {
			panic("Tenant not found")
		}

		err := database.SetTenantSchema(database.DB, tenant.Name)
		if err != nil {
			panic(err)
		}
		ctx := context.WithValue(c.Request.Context(), tenantKey{}, tenant.Name)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
