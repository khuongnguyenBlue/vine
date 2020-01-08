package exam

import "github.com/khuongnguyenBlue/vine/models"

type Service interface {
	FetchBySubjectID(subjectID uint) ([]models.Exam, error)
	GetByID(id uint) (models.Exam, error)
	GetByIDWithQuestionsAnswers(id uint) (models.Exam, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FetchBySubjectID(subjectID uint) ([]models.Exam, error) {
	return s.Repo.FetchBySubjectID(subjectID)
}

func (s *service) GetByID(id uint) (models.Exam, error) {
	return s.Repo.GetByID(id)
}

func (s *service) GetByIDWithQuestionsAnswers(id uint) (models.Exam, error) {
	return s.Repo.GetByIDWithQuestionsAnswers(id)
}
