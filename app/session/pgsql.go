package session

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
	"log"
)

type repository struct {
	Conn *gorm.DB
}

func NewPgsqlRepository(conn *gorm.DB) Repository {
	return &repository{Conn: conn}
}

func (r *repository) FindByPhoneNumber(phoneNumber string) (models.User, error) {
	var user models.User
	err := r.Conn.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *repository) Create(phoneNumber, encryptedPassword, username string) error {
	user := models.User{
		PhoneNumber: phoneNumber,
		EncryptedPassword: encryptedPassword,
		Name: username,
	}

	if err := r.Conn.Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
