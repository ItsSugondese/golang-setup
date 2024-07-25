package user

import (
	"net/http"
	"wabustock/pkg/utils"
	dto_utils "wabustock/pkg/utils/dto-utils"

	"github.com/google/uuid"
)

func SaveBaseUserService(dto UserRequest) BaseUser {
	//attachment := temporary_attachments.FindByIdService(dto.FileId)

	var response BaseUser

	if dto.ID != uuid.Nil {
		response = FindUserByIdService(dto.ID)
	}

	dto_utils.NullAwareMapDtoConvertor(dto, &response)

	return response
}

func GetUserImageService(id uuid.UUID, w http.ResponseWriter) {
	userDetails := FindUserByIdService(id)

	utils.GetFileFromFilePath(userDetails.ProfilePath, w)

}

func FindUserByIdService(id uuid.UUID) BaseUser {

	userDetails, err := FindUserByIdRepo(id)
	if userDetails.ID == uuid.Nil && err == nil {
		panic("User with that Id not found")
	}
	return userDetails

}
