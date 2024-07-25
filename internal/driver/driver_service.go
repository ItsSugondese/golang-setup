package driver

import (
	filepathconstants "easy-rider/constants/file_path_constants"
	"easy-rider/internal/user"
	"easy-rider/pkg/common/database"
	"easy-rider/pkg/utils"
)

func CreateDriverService(driverDto DriverDto) (Driver, error) {
	// Now baseUser acts as details for a driver
	var baseUser user.BaseUser
	baseUser.Name = driverDto.Name
	baseUser.Email = driverDto.Email
	baseUser.PhoneNumber = driverDto.PhoneNumber
	baseUser.Age = driverDto.Age
	baseUser.Gender = driverDto.Gender

	//hash the password using bycrypt
	hashedPassword, err := utils.HashPassword(driverDto.Password)
	if err != nil {
		return Driver{}, err

	}
	baseUser.Password = hashedPassword
	//saving image
	imageName, err := utils.SaveFile(driverDto.Image, filepathconstants.UserFilePath)
	if err != nil {
		return Driver{}, err
	}
	baseUser.ImageName = imageName

	baseUser.Role = "driver"

	// Calling repo to save base user_type_constants
	details, err := user.SaveBaseUser(baseUser)
	if err != nil {
		return Driver{}, err
	}

	driver := Driver{
		UserID:  details.ID, // Foreign key
		Details: details,
	}
	return SaveDriverRepo(driver)
}

func UpdateDriverService(driverDto DriverDto, id string) (Driver, error) {
	// Finding driver by id
	var driverFromDB Driver
	if err := database.DB.First(&driverFromDB, id).Error; err != nil {
		return Driver{}, err
	}

	//finding base user_type_constants by id
	var baseUser user.BaseUser
	if err := database.DB.First(&baseUser, driverFromDB.UserID).Error; err != nil {
		return Driver{}, err
	}

	// Update driver fields

	baseUser.Name = driverDto.Name
	baseUser.Email = driverDto.Email
	baseUser.PhoneNumber = driverDto.PhoneNumber

	baseUser.Age = driverDto.Age
	baseUser.Gender = driverDto.Gender

	//if password is not empty then hash the password otherwise keep the old password
	if driverDto.Password != "" {
		hashedPassword, err := utils.HashPassword(driverDto.Password)
		if err != nil {
			return Driver{}, err

		}
		baseUser.Password = hashedPassword
	}

	// SAVE the file and update the image URL
	// Check if an image file was provided
	if driverDto.Image != nil {
		// SAVE the file and update the image URL
		imageURL, err := utils.SaveFile(driverDto.Image, filepathconstants.UserFilePath)
		if err != nil {
			return Driver{}, err
		}
		baseUser.ImageName = imageURL
	}

	// Calling repo to save base user_type_constants

	details, err := user.UpdateBaseUser(baseUser)
	if err != nil {
		return Driver{}, err
	}

	//setting the details
	driverFromDB.Details = details

	return UpdateDriverRepo(driverFromDB)
}
func GetAlldriversService() ([]Driver, error) {
	return FindAllDriverRepo()
}

func GetdriverByIdService(id int) (Driver, error) {
	return FinddriverByIdRepo(id)
}

func DeleteDriverService(id int) error {
	return DeleteDriverRepo(id)
}
