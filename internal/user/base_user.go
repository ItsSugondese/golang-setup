package user

import (
	generic_models "wabustock/generics/generic-models"
)

type BaseUser struct {
	generic_models.AuditModel
	FullName           string `json:"fullName"`
	Email              string `json:"email" gorm:"unique"`
	PhoneNumber        string `json:"phoneNumber" gorm:"unique"`
	Password           string `json:"password"`
	Role               string `json:"role"`
	ProfilePath        string `json:"profilePath"`
	Address            string `json:"address"`
	Gender             string `json:"gender"`
	IsAccountNonLocked bool   `json:"isAccountNonLocked" gorm:"default:true"`
	IsEmailVerified    bool   `json:"isEmailVerified" gorm:"default:false"`
	UserType           string `json:"userType"`
}
