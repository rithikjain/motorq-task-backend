package models

type Course struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
