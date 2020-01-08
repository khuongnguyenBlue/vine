package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/dtos"
	"github.com/khuongnguyenBlue/vine/utils"
	"net/http"
)

func (ctl *Controller) GetExams(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if _, err = ctl.SubjectService.GetByID(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	exams, err := ctl.ExamService.FetchBySubjectID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	var examsDTO dtos.ExamsList
	examsDTO.Extract(exams)
	c.JSON(http.StatusOK, examsDTO)
}

func (ctl *Controller) GetExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exam, err := ctl.ExamService.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var examDTO dtos.Exam
	examDTO.Extract(exam)
	c.JSON(http.StatusOK, examDTO)
}

func (ctl *Controller) TakeExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exam, err := ctl.ExamService.GetByIDWithQuestionsAnswers(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var fullExamDTO dtos.FullExam
	fullExamDTO.Extract(exam)
	c.JSON(http.StatusOK, fullExamDTO)
}
