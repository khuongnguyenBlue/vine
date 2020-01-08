package subject

import "github.com/khuongnguyenBlue/vine/models"

type Repository interface {
	Fetch() ([]models.Subject, error)
	GetByID(id uint) (models.Subject, error)
}
