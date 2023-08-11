package models

type Role struct {
	ID          int    `gorm:"primary key;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDeleted   bool   `json:"is_deleted"`
}
