package user

import (
	"database/sql"
	generic_models "wabustock/generics/generic-models"
)

type BaseUser struct {
	generic_models.AuditModel
	FullName           sql.NullString `json:"fullName"`
	Email              sql.NullString `json:"email" gorm:"unique"`
	PhoneNumber        sql.NullString `json:"phoneNumber" gorm:"unique;not null"`
	Password           sql.NullString `json:"password"`
	Role               sql.NullString `json:"role"`
	ProfilePath        sql.NullString `json:"profilePath"`
	Address            sql.NullString `json:"address"`
	Gender             sql.NullString `json:"gender"`
	IsAccountNonLocked bool           `json:"isAccountNonLocked" gorm:"default:true"`
	IsEmailVerified    bool           `json:"isEmailVerified" gorm:"default:false"`
	UserType           sql.NullString `json:"userType"`
}
