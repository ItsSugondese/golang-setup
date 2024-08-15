package user

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
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
		response.ProfilePath = &filePath
	}

	mapstructure.Decode(dto, &response)
	//dto_utils.NullAwareMapDtoConvertor(dto, &response)

	var usr BaseUser
	var err error
	if response.ID == uuid.Nil {

		usr, err = SaveBaseUser(response)
	} else {
		usr, err = UpdateBaseUser(response)
	}

	if err != nil {
		panic("Error while saving base user: " + err.Error())
	}

	return usr
}

func GetUserImageService(id uuid.UUID, w http.ResponseWriter) {
	userDetails := FindUserByIdService(id)

	utils.GetFileFromFilePath(*userDetails.ProfilePath, w)

}

func FindUserByIdService(id uuid.UUID) BaseUser {

	userDetails, err := FindUserByIdRepo(id)
	if userDetails.ID == uuid.Nil && err == nil {
		panic("User with that Id not found")
	}
	return userDetails

}
