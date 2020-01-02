package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/configs"
)

type Subject struct {
	ID        uint
	Name      string `gorm:"type:varchar(20); unique; not null"`
	Questions []Question
}

func PreloadQuestions() *gorm.DB {
	return configs.DB.Preload("Questions")
}

func GetSubjects(db *gorm.DB, subjects *[]Subject) error {
	if err := db.Find(subjects).Error; err != nil {
		return err
	}

	return nil
}

func GetSubject(db *gorm.DB, subject *Subject) error {
	if err := db.First(subject).Error; err != nil {
		return err
	}

	return nil
}

func (s Subject) GetExams(db *gorm.DB, exams *[]Exam) error {
	if err := db.Model(&s).Preload("Questions").Related(&exams).Error; err != nil {
		return err
	}

	return nil
}
