package role

import (
	"errors"
	"gorm.io/gorm"
	"wabustock/pkg/common/database"
)

func SaveRoleRepo(role *Role) (*Role, error) {
	result := database.DB.Create(&role)
	return role, result.Error
}

func FindRoleByIdRepo(name *string) (role *Role, err error) {
	if err := database.DB.
		Where("id = ?", name).
		First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, return zero value of Staff and nil error
			return nil, nil
		}
		// Other errors occurred, return the error
		return nil, err
	}

	// Staff found, return the Staff and nil error
	return role, nil
}
