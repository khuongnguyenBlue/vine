package session

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/khuongnguyenBlue/vine/dtos"
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/pkg/e"
	"github.com/khuongnguyenBlue/vine/utils"
)

type Service interface {
	Authenticate(loginAcc dtos.LoginAccount) (models.User, bool)
	CreateAccount(registerInfo dtos.RegisterInfo) int
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) Authenticate(loginAcc dtos.LoginAccount) (models.User, bool)  {
	user, err := s.Repo.FindByPhoneNumber(loginAcc.PhoneNumber)
	if err != nil || !utils.ComparePassword(user.EncryptedPassword, loginAcc.Password) {
		return models.User{}, false
	}

	return user, true
}

func (s *service) CreateAccount(registerInfo dtos.RegisterInfo) int {
	encryptedPassword, err := utils.EncryptPassword(registerInfo.Password)
	if err != nil {
		return e.InternalServerError
	}

	username := randomdata.SillyName()
	err = s.Repo.Create(registerInfo.PhoneNumber, encryptedPassword, username)
	if err != nil {
		return e.ExistedPhoneNumber
	}

	return 0
}
