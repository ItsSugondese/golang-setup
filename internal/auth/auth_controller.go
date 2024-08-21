package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @BasePath /api/v1

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags AUth
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/login [post]
// get /user_type_constants/doc/:id
func LoginUser(c *gin.Context, validate *validator.Validate) {
	var authRequest AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validate.Struct(authRequest); err != nil {
		// If validation fails, return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authResponse, err := LoginService(authRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "type": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": authResponse})

}

func TestPayload(c *gin.Context, validate *validator.Validate) {
	var authRequest AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validate.Struct(authRequest); err != nil {
		// If validation fails, return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authResponse, err := LoginService(authRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "type": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": authResponse})

}

// VerifyToken is a handler function that verifies the provided paseto-token
//func VerifyToken(c *gin.Context) {
//	var authVerify AuthVerify
//	if err := c.ShouldBindJSON(&authVerify); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "type": "error", "status": 400})
//		return
//	}
//
//	verified, err := VerifyTokenService(authVerify)
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "type": "error", "status": 401})
//		return
//	}
//
//	if verified {
//
//		c.JSON(http.StatusOK, gin.H{"message": "verify successful", "status": 200, "type": "success"})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"message": "verify faild", "status": 400, "type": "error"})
//}
