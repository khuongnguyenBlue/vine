package database

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
	"log"
)

func SeedData(db *gorm.DB) error {
	if ssErr := seedSubjects(db); ssErr != nil {
		return ssErr
	}

	if sqErr := seedQuestions(db); sqErr != nil {
		return sqErr
	}

	if seErr := seedExams(db); seErr != nil {
		return seErr
	}

	if sqaErr := seedQuestionAnswers(db); sqaErr != nil {
		return sqaErr
	}

	return nil
}

func seedSubjects(db *gorm.DB) error {
	if !db.First(&models.Subject{}).RecordNotFound() {
		log.Println("Subject existed")
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	subjectNames := []string{"Toán", "Lý", "Hóa", "Sinh", "Anh"}
	for _, subject := range subjectNames {
		if err := tx.Create(&models.Subject{Name: subject}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func seedQuestions(db *gorm.DB) error {
	if !db.First(&models.Question{}).RecordNotFound() {
		log.Println("Question existed")
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var subjects []models.Subject
	if err := tx.Find(&subjects).Error; err != nil {
		return err
	}

	for _, subject := range subjects {
		for i := 0; i < 20; i++ {
			question := models.Question{Content: randomdata.Paragraph(), Subject: subject}
			if err := tx.Create(&question).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func seedExams(db *gorm.DB) error {
	if !db.First(&models.Exam{}).RecordNotFound() {
		log.Println("Exam existed")
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var subjects []models.Subject
	if err := tx.Preload("Questions").Find(&subjects).Error; err != nil {
		return err
	}

	sampleTimeAllow := [5]uint{15, 30, 45, 60}
	for _, subject := range subjects {
		for i := 0; i < 4; i++ {
			exam := models.Exam{
				Name:      randomdata.SillyName(),
				TimeAllow: sampleTimeAllow[i],
				Subject:   subject,
				Questions: subject.Questions[0 : (i+1)*5],
			}

			if err := tx.Create(&exam).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func seedQuestionAnswers(db *gorm.DB) error {
	if !db.First(&models.QuestionAnswer{}).RecordNotFound() {
		log.Println("Question answers existed")
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var questions []models.Question
	if err := tx.Find(&questions).Error; err != nil {
		return err
	}

	for _, question := range questions {
		randNum := randomdata.Number(0, 4)
		for i := 0; i < 4; i++ {
			questionAnswer := models.QuestionAnswer{
				Content:   randomdata.Country(0),
				IsCorrect: i == randNum,
				Question:  question,
			}

			if err := tx.Create(&questionAnswer).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
