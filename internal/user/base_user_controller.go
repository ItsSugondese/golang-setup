package user

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	response_crud_enum "wabustock/enums/interface-enums/response/response-crud-enum"
	localization_enums "wabustock/enums/struct-enums/localization-enums"
	"wabustock/enums/struct-enums/project_module"
	generic_controller "wabustock/generics/generic-controller"
	"wabustock/pkg/common/localization"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func RegisterUser(c *gin.Context, validate *validator.Validate) {
	var userDto UserRequest

	if err := generic_controller.ControllerValidationHandler(&userDto, c, validate); err != nil {
		return
	}

	path := SaveBaseUserService(userDto)
	generic_controller.GenericControllerSuccessResponseHandler(c, "User registered successfully!", path)
}

func Test(c *gin.Context) {
	//var userDto TestDto

	generic_controller.GenericControllerSuccessResponseHandler(c,
		localization.GetLocalizedMessage(c, localization_enums.MessageCodeEnums.API_OPERATION, map[string]interface{}{
			"First":  project_module.ModuleNameEnums.BASE_USER,
			"Second": response_crud_enum.Create(),
		}), "hello there")
}

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags user_type_constants
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user_type_constants/doc/:id [get]
// get /user/doc/:id
func GetUserImage(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	GetUserImageService(id, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "type": "error"})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"data": drivers})
}
