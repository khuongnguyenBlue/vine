package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/configs"
)

type Subject struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" gorm:"type:varchar(20); unique; not null"`
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
