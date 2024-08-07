package user

import (
	"database/sql"
	"github.com/google/uuid"
)

type UserRequest struct {
	ID uuid.UUID `json:"id"`
	//UserType string    `json:"userType" binding:"required,validUserType"`
	FileId   uuid.UUID      `json:"fileId"`
	FullName sql.NullString `json:"fullName" binding:"required"`
	//Email       string    `json:"email" binding:"required"`
	PhoneNumber sql.NullString `json:"phoneNumber" binding:"required"`
	Password    sql.NullString `json:"password" binding:"required"`
	//Role        string    `json:"role"`
	//Address string `json:"address" binding:"required"`
	//Gender      string    `json:"gender" binding:"required,validGenderType"`
}
