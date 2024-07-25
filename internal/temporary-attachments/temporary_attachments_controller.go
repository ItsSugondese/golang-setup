package temporary_attachments

import (
	"github.com/gin-gonic/gin"
	"net/http"
	pagination_utils "wabustock/pkg/utils/pagination-utils"
)

// @Summary Get list of attachment id
// @Schemes
// @Description do ping
// @Tags Temporary Attachments
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} int
// @Router /temporary-attachments [post]
// post /temporary-attachments
func CreateTemporaryAttachments(c *gin.Context) {
	var attachmentsDetailRequest TemporaryAttachmentsDetailRequest
	if err := c.ShouldBind(&attachmentsDetailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the uploaded file

	// calling service

	attach := SaveTemporaryAttachmentsService(c)

	c.JSON(http.StatusCreated, gin.H{"message": "create success", "data": attach, "type": "success"})
}

func GetTemporaryAttachments(c *gin.Context) {
	var attachmentsDetailRequest pagination_utils.PaginationRequest
	if err := c.ShouldBind(&attachmentsDetailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create success", "data": GetTempAttachmentService(attachmentsDetailRequest), "type": "success"})
}
