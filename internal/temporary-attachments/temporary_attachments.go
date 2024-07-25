package temporary_attachments

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wabustock/constants/file_type_constants"
)

type TemporaryAttachments struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`

	Name     string                       `json:"name"`
	Location string                       `json:"location"`
	FileSize float64                      `json:"file_size"`
	FileType file_type_constants.FileType `json:"file_type"`
}

type AttachmentsGorm struct {
	db *gorm.DB
}
