package role

type Role struct {
	ID *string `json:"name" gorm:"primarykey"`
}
