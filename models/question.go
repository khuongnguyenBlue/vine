package models

import "github.com/jinzhu/gorm"

type Question struct {
	ID        uint
	Content   string `gorm:"type:varchar(500); not null"`
	SubjectID uint   `gorm:"not null"`
	Subject   Subject
}

func GetQuestions(db *gorm.DB, questions *[]Question) error {
	if err := db.Find(questions).Error; err != nil {
		return err
	}

	return nil
}
