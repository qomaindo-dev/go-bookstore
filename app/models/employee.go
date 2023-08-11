package models

type Employee struct {
	ID        int     `gorm:"primary key;autoIncrement" json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Photo     string  `json:"photo"`
	IsActive  bool    `json:"is_active"`
	CreatedOn float64 `json:"created_on"`
	CreatedBy int     `json:"created_by"`
	UpdatedOn float64 `json:"updated_on"`
	UpdatedBy int     `json:"updated_by"`
}
