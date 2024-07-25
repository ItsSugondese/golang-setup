package temporary_attachments

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	generic_repo "wabustock/generics/generic-repo"
	"wabustock/pkg/common/database"
	dto_utils "wabustock/pkg/utils/dto-utils"
	pagination_utils "wabustock/pkg/utils/pagination-utils"
)

func SaveTemporaryAttachmentsRepo(attachment TemporaryAttachments) (TemporaryAttachments, error) {
	result := database.DB.Create(&attachment)
	return attachment, result.Error
}

func FindTempAttachmentsByIdRepo(id uuid.UUID) (attachment TemporaryAttachments, err error) {
	if err := database.DB.
		Where("id= ?", id).
		First(&attachment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, return zero value of Staff and nil error
			return TemporaryAttachments{}, nil
		}
		// Other errors occurred, return the error
		return TemporaryAttachments{}, err
	}

	// Staff found, return the Staff and nil error
	return attachment, nil
}

func (cg *AttachmentsGorm) ViewPages(pagination pagination_utils.PaginationRequest, response pagination_utils.PaginationResponse) *pagination_utils.PaginationResponse {
	var categories []TemporaryAttachments
	cg.db.Scopes(generic_repo.Paginate(categories, &pagination, &response, cg.db)).Find(&categories)
	tempDtos := dto_utils.ConvertSlice[TemporaryAttachments, Temp](categories)
	response.Data = tempDtos
	return &response
}
