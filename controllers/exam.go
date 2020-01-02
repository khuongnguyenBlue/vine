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
		var examsResponse serializers.ExamsList
		for _, exam := range exams {
			examsResponse.Exams = append(examsResponse.Exams, serializers.Exam{
				ID:             exam.ID,
				Name:           exam.Name,
				TimeAllow:      exam.TimeAllow,
				Status:         exam.Status,
				QuestionsCount: len(exam.Questions),
				SubjectID:      uint(id),
			})
		}
		c.JSON(http.StatusOK, examsResponse)
	}
}
