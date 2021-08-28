package models

import "time"

type Class struct {
	ID           string     `gorm:"default:uuid_generate_v4();primary_key;"`
	CourseCode   string     `json:"course_code"`
	Faculty      string     `json:"faculty"`
	Day          int        `json:"day"`
	BuildingName string     `json:"building_name"`
	Building     Building   `json:"building" gorm:"foreignKey:BuildingName"`
	StartTime    *time.Time `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Students     []Student  `json:"students" gorm:"many2many:student_classes;"`
}
