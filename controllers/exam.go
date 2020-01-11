package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/dtos"
	"github.com/khuongnguyenBlue/vine/pkg/e"
	"github.com/khuongnguyenBlue/vine/utils"
	"net/http"
)

func (ctl *Controller) GetExams(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	if _, err = ctl.SubjectService.GetByID(id); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, e.NewError(e.NotFound))
		return
	}

	exams, err := ctl.ExamService.FetchBySubjectID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var examsDTO dtos.ExamsList
	examsDTO.Extract(exams)
	c.JSON(http.StatusOK, examsDTO)
}

func (ctl *Controller) GetExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	exam, err := ctl.ExamService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, e.NewError(e.NotFound))
		return
	}

	var examDTO dtos.Exam
	examDTO.Extract(exam)
	c.JSON(http.StatusOK, examDTO)
}

func (ctl *Controller) TakeExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	exam, err := ctl.ExamService.GetByIDWithQuestionsAnswers(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, e.NewError(e.NotFound))
		return
	}

	var fullExamDTO dtos.FullExam
	fullExamDTO.Extract(exam)
	c.JSON(http.StatusOK, fullExamDTO)
}

func (ctl *Controller) SubmitExam(c *gin.Context) {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	var submittedExam dtos.SubmittedExam
	if err := c.ShouldBindJSON(&submittedExam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	if id != submittedExam.ID {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	userID, ok := utils.GetUserID(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	examResult, operationCode := ctl.ExamService.SaveSubmittedExam(submittedExam, userID)
	switch operationCode {
	case http.StatusInternalServerError:
		c.AbortWithStatus(http.StatusInternalServerError)
	case http.StatusBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
	case http.StatusNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, e.NewError(e.NotFound))
	default:
		var briefResultDTO dtos.BriefResult
		briefResultDTO.Extract(examResult)
		c.JSON(http.StatusOK, briefResultDTO)
	}
}

func (ctl *Controller) ReviewExam(c *gin.Context)  {
	id, err := utils.GetIDParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.NewError(e.InvalidParams))
		return
	}

	userID, ok := utils.GetUserID(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	reviewExam, operationCode := ctl.ExamService.GetExamForReview(id, userID)
	switch operationCode {
	case http.StatusNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, e.NewError(e.NotFound))
	default:
		c.JSON(http.StatusOK, reviewExam)
	}
}
