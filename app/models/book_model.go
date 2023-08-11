package models

type Book struct {
	ID        int     `gorm:"primary key;autoIncrement" json:"id"`
	Title     string  `json:"title"`
	Synopsis  string  `json:"synopsis"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Stock     int     `json:"int"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
	CreatedOn float64 `json:"created_on"`
	CreatedBy int     `json:"created_by"`
	UpdatedOn float64 `json:"updated_on"`
	UpdatedBy int     `json:"updated_by"`
}
