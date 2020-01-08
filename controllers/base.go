package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/app/session"
	"github.com/khuongnguyenBlue/vine/app/exam"
	"github.com/khuongnguyenBlue/vine/app/subject"
)

type Controller struct {
	SessionService session.Service
	SubjectService subject.Service
	ExamService    exam.Service
}

func NewController(db *gorm.DB) *Controller {
	sessionService := session.NewService(session.NewPgsqlRepository(db))
	subjectService := subject.NewService(subject.NewPgsqlRepository(db))
	examService := exam.NewService(exam.NewPgsqlRepository(db))

	return &Controller{
		SessionService: sessionService,
		SubjectService: subjectService,
		ExamService:    examService,
	}
}
