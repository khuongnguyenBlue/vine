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
	} else {
		var subjectsResponse serializers.SubjectsList
		for _, subject := range subjects {
			subjectsResponse.Subjects = append(subjectsResponse.Subjects, serializers.Subject{
				ID:   subject.ID,
				Name: subject.Name,
			})
		}
		c.JSON(http.StatusOK, subjectsResponse)
	}
}
