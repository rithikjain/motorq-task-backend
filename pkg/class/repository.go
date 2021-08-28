package class

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllCourses() (*[]models.Course, error)
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
