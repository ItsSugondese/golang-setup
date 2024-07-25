package driver

import (
	"MahabiraLogistics/internal/user"
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Driver struct {
	UserID  uuid.UUID     `json:"user_id"`                          // Foreign key
	Details user.BaseUser `json:"details" gorm:"foreignKey:UserID"` // One-to-one relationship

}
