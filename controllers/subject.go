package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/dtos"
	"net/http"
)

func (ctl *Controller) GetSubjects(c *gin.Context) {
	subjects, err := ctl.SubjectService.Fetch()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var subjectsDTO dtos.SubjectsList
	subjectsDTO.Extract(subjects)
	c.JSON(http.StatusOK, subjectsDTO)
}
