package subject

import (
	"github.com/khuongnguyenBlue/vine/models"
)

type Service interface {
	Fetch() ([]models.Subject, error)
	GetByID(id uint) (models.Subject, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) Fetch() ([]models.Subject, error)  {
	return s.Repo.Fetch()
}

func (s *service) GetByID(id uint) (models.Subject, error) {
	return s.Repo.GetByID(id)
}
