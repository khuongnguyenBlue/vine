package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/app/auth"
	"github.com/khuongnguyenBlue/vine/app/exam"
	"github.com/khuongnguyenBlue/vine/app/subject"
)

type Controller struct {
	AuthService auth.Service
	SubjectService subject.Service
	ExamService exam.Service
}

func NewController(db *gorm.DB) *Controller {
	authService := auth.NewService(auth.NewPgsqlRepository(db))
	subjectService := subject.NewService(subject.NewPgsqlRepository(db))
	examService := exam.NewService(exam.NewPgsqlRepository(db))

	return &Controller{
		AuthService: authService,
		SubjectService: subjectService,
		ExamService: examService,
	}
}
