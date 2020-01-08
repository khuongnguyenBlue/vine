package auth

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

func (r *repository) FindByPhoneNumber(phoneNumer string) (models.User, error) {
	user := models.User{PhoneNumber: phoneNumer}
	err := r.Conn.First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
