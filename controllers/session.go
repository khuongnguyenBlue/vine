package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/dtos"
	"github.com/khuongnguyenBlue/vine/pkg/e"
	"github.com/khuongnguyenBlue/vine/utils"
	"net/http"
)

func (ctl *Controller) Login(c *gin.Context) {
	var loginAcc dtos.LoginAccount
	if err := c.ShouldBind(&loginAcc); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidLoginAccount))
		return
	}

	user, valid := ctl.SessionService.Authenticate(loginAcc)
	if !valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, e.NewError(e.WrongLoginAccount))
		return
	}

	token, err := utils.GenerateToken(user.Name, user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctl *Controller) Register(c *gin.Context) {
	var registerInfo dtos.RegisterInfo
	if err := c.ShouldBind(&registerInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidRegisterAccount))
		return
	}

	switch errCode := ctl.SessionService.CreateAccount(registerInfo) ; errCode {
	case http.StatusInternalServerError:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	case e.ExistedPhoneNumber:
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(errCode))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
