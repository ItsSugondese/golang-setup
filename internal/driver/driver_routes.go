package driver

import (
	authentication_middleware "easy-rider/pkg/middleware/authentication-middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DriverRoutes(r *gin.Engine, validate *validator.Validate) {
	drivers := r.Group("/driver/")
	{
		drivers.POST("create", CreateDriver)
		drivers.Use(authentication_middleware.AuthMiddleware())
		drivers.PUT("update/:id", UpdateDriver)
		drivers.GET("get-all", GetDrivers)
		drivers.GET("get/:id", GetDriverById)
		drivers.DELETE("delete/:id", DeleteDriver)
	}
}
