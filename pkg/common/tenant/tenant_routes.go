package tenant

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TenantRoutes(r *gin.Engine, validate *validator.Validate) {
	tenants := r.Group("/tenant/")
	{
		tenants.POST("", func(c *gin.Context) {
			CreateTenant(c, validate)
		})
	}
}
