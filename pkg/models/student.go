package models

type Student struct {
	RollNo  string  `json:"roll_no" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Classes []Class `json:"classes" gorm:"many2many:student_classes;"`
}
