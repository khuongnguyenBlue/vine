package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"github.com/khuongnguyenBlue/vine/serializers"
	"github.com/khuongnguyenBlue/vine/utils"
	"net/http"
)

func GetExams(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	subject := models.Subject{ID: id}
	if err = models.GetSubject(configs.DB, &subject); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var exams []models.Exam
	if err = subject.GetExams(configs.DB, &exams); err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	var examsJson serializers.ExamsJson
	examsJson.Parse(exams)
	c.JSON(http.StatusOK, examsJson)
}

func GetExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exam := models.Exam{ID: id}
	if err = models.GetExam(models.PreloadQuestions(), &exam); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var examJson serializers.ExamJson
	examJson.Parse(exam)
	c.JSON(http.StatusOK, examJson)
}
