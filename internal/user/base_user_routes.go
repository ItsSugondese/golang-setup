package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	authentication_middleware "wabustock/pkg/middleware/authentication-middleware"
	"wabustock/pkg/utils/paseto-token"
)

func UserRoutes(r *gin.Engine, validate *validator.Validate) {
	users := r.Group("/user/")
	{
		users.POST("", func(c *gin.Context) {
			RegisterUser(c, validate)
		})
		users.POST("test", Test)
		//users.Use(authentication_middleware.AuthMiddleware())
		users.GET("doc/:id", GetUserImage)
		users.Use(authentication_middleware.PasetoAuthMiddleware(*paseto_token.TokenMaker))
	}
}
