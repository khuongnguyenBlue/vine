package controllers

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/utils"
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user models.User
	encryptedPassword, _ := utils.EncryptPassword(password)

	// tạm thời tạo mới tài khoản nếu nhập SĐT mới
	err := configs.DB.Where(models.User{PhoneNumber: phoneNumber}).
		Attrs(models.User{Name: randomdata.SillyName(), EncryptedPassword: encryptedPassword}).
		FirstOrInit(&user).Error

	if err != nil {
		log.Println("Something's wrong")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		configs.DB.Create(&user)
	} else if !utils.ComparePassword(user.EncryptedPassword, password) {
		log.Println("Wrong password")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Name, user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
