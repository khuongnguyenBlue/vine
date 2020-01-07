package exam

import "github.com/khuongnguyenBlue/vine/models"

type Repository interface {
	FetchBySubjectID(subjectId uint) ([]models.Exam, error)
	GetByID(id uint) (models.Exam, error)
}