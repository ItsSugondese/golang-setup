package driver

import (
	"mime/multipart"
)

type DriverDto struct {
	Name        string                `form:"name" binding:"required"`
	Email       string                `form:"email" binding:"required"`
	PhoneNumber string                `form:"phoneNumber" binding:"required"`
	Password    string                `form:"password" binding:"required"`
	Age         int                   `form:"age" binding:"required"`
	Gender      string                `form:"gender" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	// Image *multipart.FileHeader `form:"image" `
}
