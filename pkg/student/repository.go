package student

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type Repository interface {
	CreateStudent(student *models.Student) error
	GetStudentByRollNo(rollNo string) (*models.Student, error)
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

func (r *repo) GetStudentByRollNo(rollNo string) (*models.Student, error) {
	student := &models.Student{}
	err := r.DB.Where("LOWER(roll_no)=?", strings.ToLower(rollNo)).First(student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound
		} else {
			return nil, utils.ErrDatabase
		}
	}
	return student, nil
}
