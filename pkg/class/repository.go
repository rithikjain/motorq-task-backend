package class

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllCourses() (*[]models.Course, error)
	GetAllClassesForACourse(courseID string) (*[]models.Class, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) GetAllCourses() (*[]models.Course, error) {
	var courses []models.Course
	err := r.DB.Find(&courses).Error
	if err != nil {
		return nil, utils.ErrDatabase
	}
	return &courses, nil
}

func (r *repo) GetAllClassesForACourse(courseID string) (*[]models.Class, error) {
	var classes []models.Class
	err := r.DB.Where("course_id=?", courseID).Preload("Building").Preload("Course").Find(&classes).Error
	if err != nil {
		return nil, utils.ErrDatabase
	}
	return &classes, nil
}
