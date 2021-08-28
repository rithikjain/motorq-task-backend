package models

import "time"

type Class struct {
	ID                 string     `gorm:"default:uuid_generate_v4();primary_key;"`
	CourseID           string     `json:"-"`
	Course             Course     `json:"course,omitempty"`
	Faculty            string     `json:"faculty"`
	Day                int        `json:"day"`
	BuildingName       string     `json:"-"`
	Building           Building   `json:"building" gorm:"foreignKey:BuildingName"`
	StartTime          *time.Time `json:"start_time"`
	EndTime            *time.Time `json:"end_time"`
	TimeString         string     `json:"time_string" gorm:"-"`
	Students           []Student  `json:"-" gorm:"many2many:student_classes;"`
	StudentsRegistered int64      `json:"students_registered" gorm:"-"`
}
