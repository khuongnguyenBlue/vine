package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/serializers"
	"net/http"
)

func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	err := models.GetSubjects(configs.DB, &subjects)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	var subjectsJson serializers.SubjectsJson
	subjectsJson.Parse(subjects)
	c.JSON(http.StatusOK, subjectsJson)
}
