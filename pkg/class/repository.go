package class

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	GetAllCourses() (*[]models.Course, error)
	GetAllClassesForACourse(courseID string) (*[]models.Class, error)
	AddClassStudent(studentID, classID string) error
	GetEnrolledClasses(studentID string) (*[]models.Class, error)
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
	rows, err := r.DB.Model(&models.Class{}).Where("course_id=?", courseID).Rows()
	if err != nil {
		return nil, utils.ErrDatabase
	}

	defer rows.Close()
	for rows.Next() {
		var class models.Class
		err = r.DB.ScanRows(rows, &class)
		if err != nil {
			return nil, utils.ErrDatabase
		}

		err = r.DB.Model(&class).Association("Building").Find(&class.Building)
		if err != nil {
			return nil, utils.ErrDatabase
		}

		err = r.DB.Model(&class).Association("Course").Find(&class.Course)
		if err != nil {
			return nil, utils.ErrDatabase
		}

		count := r.DB.Model(&class).Association("StudentsRegistered").Count()
		class.StudentsRegistered = count
		classes = append(classes, class)
	}

	return &classes, nil
}

func (r *repo) AddClassStudent(studentID, classID string) error {
	var student models.Student
	err := r.DB.Where("LOWER(roll_no)=?", strings.ToLower(studentID)).First(&student).Error
	if err != nil {
		return utils.ErrDatabase
	}
	err = r.DB.Model(&student).Association("Classes").Append(&models.Class{ID: classID})
	if err != nil {
		return utils.ErrDatabase
	}

	return nil
}

func (r *repo) GetEnrolledClasses(studentID string) (*[]models.Class, error) {
	var classes []models.Class
	err := r.DB.Model(&models.Student{RollNo: studentID}).Preload("Course").Preload("Building").Association("Classes").Find(&classes)
	if err != nil {
		return nil, utils.ErrDatabase
	}
	return &classes, nil
}
