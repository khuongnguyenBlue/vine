package subject

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
)

type repository struct {
	Conn *gorm.DB
}

func NewPgsqlRepository(conn *gorm.DB) Repository {
	return &repository{Conn: conn}
}

func (r *repository) Fetch() ([]models.Subject, error) {
	var subjects []models.Subject
	if err := r.Conn.Find(&subjects).Error; err != nil {
		return nil, err
	}

	return subjects, nil
}

func (r *repository) GetByID(id uint) (models.Subject, error) {
	var subject = models.Subject{ID: id}
	if err := r.Conn.First(&subject).Error; err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}
