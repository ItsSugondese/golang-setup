package user

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"wabustock/enums/struct-enums/project_module"
	"wabustock/internal/role"
	temporary_attachments "wabustock/internal/temporary-attachments"
	"wabustock/pkg/utils"
)

func SaveBaseUserService(dto UserRequest) BaseUser {

	var response BaseUser
	mapstructure.Decode(dto, &response)
	if dto.ID != uuid.Nil {
		response = FindUserByIdService(dto.ID)
	} else {
		if dto.Role == nil {
			panic("Specifying role is must.")
		}

		getRole := role.FindRoleByIdService(dto.Role)
		roles := []role.Role{getRole}
		response.Roles = roles
	}
	if dto.FileId != uuid.Nil {
		attachment := temporary_attachments.FindByIdService(dto.FileId)
		//filePath := utils.CopyFileToServer(attachment.Location, project_module.ModuleNameEnums.BASE_USER)
		filePath := utils.CopyFileToServer(attachment.Location, project_module.ModuleNameEnums.BASE_USER, false)
		response.ProfilePath = &filePath
	}

	if dto.Password != nil {
		password, hashingPasswordError := utils.HashPassword(*dto.Password)
		if hashingPasswordError != nil {
			panic("error hashing password")
		}
		response.Password = &password
	}

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
	//userDetails := FindUserByIdService(id)
	userDetails := temporary_attachments.FindByIdService(id)
	//utils.GetFileFromFilePath(*userDetails.ProfilePath, w)
	utils.GetFileFromFilePath(userDetails.Location, w, true)

}

func FindUserByIdService(id uuid.UUID) BaseUser {

	userDetails, err := FindUserByIdRepo(id)
	if userDetails.ID == uuid.Nil && err == nil {
		panic("User with that Id not found")
	}
	return userDetails

}
