package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"net/http"
)

func GetSubjects(c *gin.Context)  {
	var subjects []models.Subject
	err := models.GetSubjects(configs.DB, &subjects)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	} else {
		c.JSON(http.StatusOK, subjects)
	}
}
