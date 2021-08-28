package class

import (
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"time"
)

type Service interface {
	GetAllCourses() (*[]models.Course, error)
	GetAllClassesForACourse(courseID string) (*[]models.Class, error)
	AddClassStudent(studentID, classID string) error
	GetEnrolledClasses(studentID string) (*[]models.Class, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetAllCourses() (*[]models.Course, error) {
	return s.repo.GetAllCourses()
}

func (s *service) GetAllClassesForACourse(courseID string) (*[]models.Class, error) {
	classes, err := s.repo.GetAllClassesForACourse(courseID)
	if err != nil {
		return nil, err
	}

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(*classes); i++ {
		timeStr := days[(*classes)[i].Day-1] + ", " + (*classes)[i].StartTime.Format(time.Kitchen) + " to " + (*classes)[i].EndTime.Format(time.Kitchen)
		(*classes)[i].TimeString = timeStr
	}
	return classes, nil
}

func (s *service) AddClassStudent(studentID, classID string) error {
	return s.repo.AddClassStudent(studentID, classID)
}

func (s *service) GetEnrolledClasses(studentID string) (*[]models.Class, error) {
	classes, err := s.repo.GetEnrolledClasses(studentID)
	if err != nil {
		return nil, err
	}

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(*classes); i++ {
		timeStr := days[(*classes)[i].Day-1] + ", " + (*classes)[i].StartTime.Format(time.Kitchen) + " to " + (*classes)[i].EndTime.Format(time.Kitchen)
		(*classes)[i].TimeString = timeStr
	}
	return classes, nil
}
