package user

import (
	"database/sql"
	"github.com/google/uuid"
	"net/http"
	"wabustock/enums/struct-enums/project_module"
	temporary_attachments "wabustock/internal/temporary-attachments"
	"wabustock/pkg/utils"
)

func SaveBaseUserService(dto UserRequest) BaseUser {

	var response BaseUser
	if dto.ID != uuid.Nil {
		response = FindUserByIdService(dto.ID)
	}

	response.PhoneNumber = dto.PhoneNumber
	if dto.FileId != uuid.Nil {
		attachment := temporary_attachments.FindByIdService(dto.FileId)
		filePath := utils.CopyFileToServer(attachment.Location, project_module.ModuleNameEnums.BASE_USER)
		response.ProfilePath = sql.NullString{
			String: filePath,
			Valid:  filePath != "",
		}
	}

	usr, err := SaveBaseUser(response)

	if err != nil {
		panic("Error while saving base user: " + err.Error())
	}
	//dto_utils.NullAwareMapDtoConvertor(dto, &response)

	return usr
}

func GetUserImageService(id uuid.UUID, w http.ResponseWriter) {
	userDetails := FindUserByIdService(id)

	utils.GetFileFromFilePath(userDetails.ProfilePath.String, w)

}

func FindUserByIdService(id uuid.UUID) BaseUser {

	userDetails, err := FindUserByIdRepo(id)
	if userDetails.ID == uuid.Nil && err == nil {
		panic("User with that Id not found")
	}
	return userDetails

}
