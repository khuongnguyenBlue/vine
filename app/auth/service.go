package auth

import (
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/utils"
)

type Service interface {
	ValidateLoginAccount(phoneNumber, password string) (models.User, bool)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) ValidateLoginAccount(phoneNumber, password string) (models.User, bool)  {
	user, err := s.Repo.FindByPhoneNumber(phoneNumber)
	if err != nil || !utils.ComparePassword(user.EncryptedPassword, password) {
		return models.User{}, false
	}

	return user, true
}
