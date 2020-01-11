package database

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
)

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		models.User{},
		models.Subject{},
		models.Question{},
		models.Exam{},
		models.QuestionAnswer{},
		models.ExamResult{},
		models.UserAnswer{},
	)
	db.Model(&models.Exam{}).AddIndex("idx_esubject_id", "subject_id")
	db.Model(&models.Question{}).AddIndex("idx_qsubject_id", "subject_id")
}
