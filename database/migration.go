package database

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		models.User{},
		models.Subject{},
		models.Question{},
		models.Exam{},
		models.QuestionAnswer{},
	)
}
