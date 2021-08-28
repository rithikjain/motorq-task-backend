package student

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateStudent(student *models.Student) error
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateStudent(student *models.Student) error {
	err := r.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "roll_no"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": student.Name}),
	}).Create(student).Error
	if err != nil {
		return utils.ErrDatabase
	}
	return nil
}
