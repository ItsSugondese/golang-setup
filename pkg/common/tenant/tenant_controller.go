package tenant

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	response_crud_enum "wabustock/enums/interface-enums/response/response-crud-enum"
	localization_enums "wabustock/enums/struct-enums/localization-enums"
	"wabustock/enums/struct-enums/project_module"
	generic_controller "wabustock/generics/generic-controller"
	"wabustock/pkg/common/localization"
)

// @Summary Create Tenant
// @Schemes
// @Description do ping
// @Tags Temporary Attachments
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} int
// @Router /tenant [post]
// post /tenant
func CreateTenant(c *gin.Context, validate *validator.Validate) {
	var tenantRequest TenantRequest

	// validate payload
	generic_controller.ControllerValidationHandler(&tenantRequest, c, validate)

	// Get from service
	response := CreateTenantSchemaService(tenantRequest)

	//response body
	generic_controller.GenericControllerSuccessResponseHandler(c,
		localization.GetLocalizedMessage(localization_enums.MessageCodeEnums.API_OPERATION, map[string]interface{}{
			"First":  project_module.ModuleNameEnums.TENANT,
			"Second": response_crud_enum.Create(),
		}), response)
}
