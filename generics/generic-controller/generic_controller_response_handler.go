package generic_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wabustock/enums/interface-enums/response/response-status-enum"
	globaldto "wabustock/global/global_dto"
)

func GenericControllerSuccessResponseHandler(c *gin.Context, message string, data interface{}) {
	response := globaldto.ApiResponse{
		Status:  response_status_enum.Success(),
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
