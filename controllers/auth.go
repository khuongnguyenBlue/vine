package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/utils"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

type account struct {
	PhoneNumber string `validate:"required,lte=11"`
	Password    string `validate:"required,lte=20"`
}

func (ctl *Controller) Login(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")
	auth := account{PhoneNumber: phoneNumber, Password: password}

	var validate = validator.New()
	err := validate.Struct(auth)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Field())
			log.Println(err.Type())
			log.Println()
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, valid := ctl.AuthService.ValidateLoginAccount(phoneNumber, password)
	if !valid {
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
