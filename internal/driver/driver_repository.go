package driver

import (
	"MahabiraLogistics/pkg/common/database"
	"errors"

	"gorm.io/gorm"
)

//func SaveDriverRepo(driver Driver) (Driver, error) {
//    result := database.DB.Create(&driver)
//    return driver, result.Error
//}
//
//func UpdateDriverRepo(driver Driver) (Driver, error) {
//
//	result := database.DB.Model(&driver).
//		Updates(driver)
//
//	 return driver, result.Error
//}

func FindAllDriverRepo() ([]Driver, error) {
	var drivers []Driver
	result := database.DB.Find(&drivers)
	//set details also in result before returning
	for i, driver := range drivers {
		database.DB.First(&drivers[i].Details, driver.UserID)

	}

	return drivers, result.Error
}

func FinddriverByIdRepo(id int) (Driver, error) {
	var driver Driver
	result := database.DB.First(&driver, id)
	//set details also in result before returning
	database.DB.First(&driver.Details, driver.UserID)
	return driver, result.Error
}

func DeleteDriverRepo(id int) error {
	var driver Driver
	result := database.DB.Delete(&driver, id)
	return result.Error
}
func FindDriverByPhoneNumberRepo(phoneNumber string) (driver Driver, err error) {
	if err := database.DB.
		Joins("JOIN base_users ON base_users.id = drivers.user_id").
		Where("base_users.phone_number = ?", phoneNumber).
		Preload("Details").
		First(&driver).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Record not found, return zero value of driver and nil error
			return Driver{}, nil
		}
		// Other errors occurred, return the error
		return Driver{}, err
	}

	// driver found, return the driver and nil error
	return driver, nil
}
