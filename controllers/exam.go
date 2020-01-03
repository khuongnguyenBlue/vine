package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/serializers"
	"net/http"
	"strconv"
)

func GetExams(c *gin.Context)  {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	subject := models.Subject{ID: uint(id)}
	if err = models.GetSubject(configs.DB, &subject); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var exams []models.Exam
	if err = subject.GetExams(configs.DB, &exams); err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	} else {
		var examsJson serializers.ExamsJson
		examsJson.Parse(exams)
		c.JSON(http.StatusOK, examsJson)
	}
}
