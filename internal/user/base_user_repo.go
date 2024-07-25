package user

import (
	"errors"
	"wabustock/pkg/common/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SaveBaseUser(user BaseUser) (BaseUser, error) {
	result := database.DB.Create(&user)
	return user, result.Error
}
func UpdateBaseUser(user BaseUser) (BaseUser, error) {
	result := database.DB.Model(&user).
		Updates(user)
	return user, result.Error
}

func FindUserByPhoneNumberRepo(phoneNumber string) (user BaseUser, err error) {
	if err := database.DB.
		Where("phone_number = ?", phoneNumber).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, return zero value of Staff and nil error
			return BaseUser{}, nil
		}
		// Other errors occurred, return the error
		return BaseUser{}, err
	}

	// Staff found, return the Staff and nil error
	return user, nil
}

func FindUserByIdRepo(id uuid.UUID) (user BaseUser, err error) {
	if err := database.DB.
		Where("id= ?", id).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, return zero value of Staff and nil error
			return BaseUser{}, nil
		}
		// Other errors occurred, return the error
		return BaseUser{}, err
	}

	// Staff found, return the Staff and nil error
	return user, nil
}
