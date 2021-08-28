package class

import "github.com/rithikjain/motorq-task-backend/pkg/models"

type Service interface {
	GetAllCourses() (*[]models.Course, error)
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
