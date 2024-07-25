package generic_models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AuditModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsActive  bool           `gorm:"default:true"`
}
