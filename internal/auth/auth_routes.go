package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AuthRoutes(r *gin.Engine, validate *validator.Validate) {
	auths := r.Group("/auth/")
	{
		auths.POST("login", func(c *gin.Context) {
			LoginUser(c, validate)
		})
		//auths.POST("verify", VerifyToken)

	}
}
