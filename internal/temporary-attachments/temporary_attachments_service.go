package temporary_attachments

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"mime/multipart"
	"wabustock/enums/struct-enums/project_module"
	"wabustock/pkg/common/database"
	"wabustock/pkg/utils"
	pagination_utils "wabustock/pkg/utils/pagination-utils"
)

func SaveTemporaryAttachmentsService(c *gin.Context) []uuid.UUID {

	var ids []uuid.UUID
	form, err := c.MultipartForm()
	if err != nil {
		panic("Error parsing form: " + err.Error())

	}

	files := form.File["attachments"] // "attachments" is the form field name for the files

	var attachments []*multipart.FileHeader

	// Loop over the files
	for _, fileHeader := range files {
		attachments = append(attachments, fileHeader)
		fileDetails := utils.SaveFile(fileHeader, project_module.ModuleNameEnums.TEMPORARY_ATTACHMENTS, true)

		attach, err := SaveTemporaryAttachmentsRepo(TemporaryAttachments{
			Name:     fileHeader.Filename,
			Location: fileDetails.FilePath,
			FileSize: math.Round(float64((fileDetails.Size/1000)*100)) / 100,
			FileType: fileDetails.FileType,
		})

		if err != nil {
			panic("Error when saving file")
		}

		ids = append(ids, attach.ID)
	}
	return ids
}

func FindByIdService(id uuid.UUID) TemporaryAttachments {
	attachment, err := FindTempAttachmentsByIdRepo(id)
	if err != nil {
		panic("Didn't find attachment with that id")
	}

	return attachment
}

func GetTempAttachmentService(request pagination_utils.PaginationRequest) *pagination_utils.PaginationResponse {
	attachmentsGorm := &AttachmentsGorm{db: database.DB}

	// Call the ViewPages function
	return attachmentsGorm.ViewPages(request, pagination_utils.PaginationResponse{})
}
