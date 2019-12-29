package controllers

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

type auth struct {
	PhoneNumber string `validate:"required,lte=11"`
	Password    string `validate:"required,lte=20"`
}

func Login(c *gin.Context) {
	var validate = validator.New()

	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")

	auth := auth{PhoneNumber: phoneNumber, Password: password}
	validationErr := validate.Struct(auth)
	if validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			log.Println(err.Field())
			log.Println(err.Type())
			log.Println()
		}
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var user models.User
	encryptedPassword, _ := encryptPassword(password)

	// tạm thời tạo mới tài khoản nếu nhập SĐT mới
	err := configs.DB.Where(models.User{PhoneNumber: phoneNumber}).
		Attrs(models.User{Name: randomdata.SillyName(), EncryptedPassword: encryptedPassword}).
		FirstOrInit(&user).Error

	if err != nil {
		log.Println("Something's wrong")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if user.ID == 0 {
		configs.DB.Create(&user)
	} else if !comparePassword(user.EncryptedPassword, password) {
		log.Println("Wrong password")
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func encryptPassword(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), 4)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func comparePassword(encryptedPassword string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(input))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
