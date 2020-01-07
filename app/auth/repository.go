package auth

import "github.com/khuongnguyenBlue/vine/models"

type Repository interface {
	FindByPhoneNumber(phoneNumber string) (models.User, error)
}
