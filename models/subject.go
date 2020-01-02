package models

import "github.com/khuongnguyenBlue/vine/configs"

type Subject struct {
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"type:varchar(20); unique; not null"`
}

func GetSubjects(subjects *[]Subject) error  {
	if err := configs.DB.Find(subjects).Error; err != nil {
		return err
	}

	return nil
}
