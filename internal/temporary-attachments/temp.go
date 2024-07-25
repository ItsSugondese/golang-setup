package temporary_attachments

import (
	"wabustock/generics/generic-models"
)

type Temp struct {
	generic_models.AuditModel
	Name     string `json:"name"`
	Location string `json:"location"`
}
