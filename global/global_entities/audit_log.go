package global_entities

import (
	"gorm.io/datatypes"
	"time"
)

type AuditLog struct {
	Id            string         `json:"id" gorm:"primaryKey"`
	TableName     string         `json:"table_name"`
	OperationType string         `json:"operation_type"`
	ObjectId      string         `json:"object_id"`
	Data          datatypes.JSON `json:"data"`
	CreatedBy     string         `json:"created_by"`
	UpdatedBy     *string        `json:"updated-by"`
	CreatedAt     time.Time      `json:"created_at"`
}
