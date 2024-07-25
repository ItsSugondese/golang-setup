package user

import (
	"github.com/go-playground/validator/v10"
	gender_type_enums "wabustock/enums/struct-enums/gender-type-enums"
	"wabustock/enums/struct-enums/user_type_enums"
)

// Define the custom validator for user_type_constants type enum defined in userRequest payload
func ValidUserType(fl validator.FieldLevel) bool {
	userType := fl.Field().String()
	return userType == user_type_enums.UserType.USER || userType == user_type_enums.UserType.COMPANY || userType == user_type_enums.UserType.DRIVER
}

func ValidGenderType(fl validator.FieldLevel) bool {
	genderType := fl.Field().String()
	return genderType == gender_type_enums.GenderType.MALE || genderType == gender_type_enums.GenderType.FEMALE
}
