package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/models"
	"net/http"
)

func GetSubjects(c *gin.Context)  {
	var subjects []models.Subject
	err := models.GetSubjects(&subjects)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subjects)
	}
}
