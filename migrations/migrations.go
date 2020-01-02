package migrations

import (
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
)

func Migrate() {
	configs.DB.AutoMigrate(
		models.User{},
		models.Subject{},
		models.Question{},
		models.Exam{},
		models.QuestionAnswer{},
	)
}
