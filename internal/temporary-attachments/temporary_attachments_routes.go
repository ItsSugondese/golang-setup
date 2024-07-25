package temporary_attachments

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TempAttachmentsRoutes(r *gin.Engine, validate *validator.Validate) {
	tempAttachent := r.Group("/temp-attachments/")
	{
		tempAttachent.POST("", func(c *gin.Context) {
			CreateTemporaryAttachments(c)
		})
		tempAttachent.POST("/get", func(c *gin.Context) {
			GetTemporaryAttachments(c)
		})
	}
}
