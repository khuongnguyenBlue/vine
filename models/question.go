package models

import "github.com/jinzhu/gorm"

type Question struct {
	ID        uint   `json:"id"`
	Content   string `json:"content" gorm:"type:varchar(500); not null"`
	SubjectID uint   `json:"subject_id" gorm:"not null"`
	Subject   Subject
}

func GetQuestions(db *gorm.DB, questions *[]Question) error {
	if err := db.Find(questions).Error; err != nil {
		return err
	}

	return nil
}
