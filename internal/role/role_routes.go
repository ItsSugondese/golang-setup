package role

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RoleRoutes(r *gin.Engine, validate *validator.Validate) {
	roles := r.Group("/role/")
	{
		roles.POST("", func(c *gin.Context) {
			CreateRole(c, validate)
		})
	}
}
