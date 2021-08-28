package student

import "github.com/rithikjain/motorq-task-backend/pkg/models"

type Service interface {
	CreateStudent(student *models.Student) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}
