package user

import (
	generic_models "wabustock/generics/generic-models"
	"wabustock/internal/role"
)

type BaseUser struct {
	generic_models.AuditModel
	FullName    *string `json:"fullName"`
	Email       *string `json:"email" gorm:"unique"`
	PhoneNumber *string `json:"phoneNumber" gorm:"unique;not null"`
	Password    *string `json:"password"`
	//Role        *string `json:"role" `
	Roles              []role.Role `json:"role" gorm:"many2many:user_role;association_autoupdate:false;association_autocreate:false"`
	ProfilePath        *string     `json:"profilePath"`
	Address            *string     `json:"address"`
	Gender             *string     `json:"gender"`
	IsAccountNonLocked *bool       `json:"isAccountNonLocked" gorm:"default:true"`
	IsEmailVerified    *bool       `json:"isEmailVerified" gorm:"default:false"`
	UserType           *string     `json:"userType"`
}

func (b *BaseUser) HasAuditModel() bool {
	return true
}
