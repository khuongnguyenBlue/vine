package seeds

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"log"
)

func All(c chan error) {
	if ssErr := seedSubjects(); ssErr != nil {
		c <- ssErr
		return
	}

	if sqErr := seedQuestions(); sqErr != nil {
		c <- sqErr
		return
	}

	if seErr := seedExams(); seErr != nil {
		c <- seErr
		return
	}

	if sqaErr := seedQuestionAnswers(); sqaErr != nil {
		c <- sqaErr
		return
	}

	c <- nil
}

func seedSubjects() error {
	if !configs.DB.First(&models.Subject{}).RecordNotFound() {
		log.Println("Subject existed")
		return nil
	}

	tx := configs.DB.Begin()
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

func seedQuestions() error {
	if !configs.DB.First(&models.Question{}).RecordNotFound() {
		log.Println("Question existed")
		return nil
	}

	tx := configs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var subjects []models.Subject
	if err := models.GetSubjects(configs.DB, &subjects); err != nil {
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

func seedExams() error {
	if !configs.DB.First(&models.Exam{}).RecordNotFound() {
		log.Println("Exam existed")
		return nil
	}

	tx := configs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var subjects []models.Subject
	if err := models.GetSubjects(models.PreloadQuestions(), &subjects); err != nil {
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

func seedQuestionAnswers() error {
	if !configs.DB.First(&models.QuestionAnswer{}).RecordNotFound() {
		log.Println("Question answers existed")
		return nil
	}

	tx := configs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var questions []models.Question
	if err := models.GetQuestions(configs.DB, &questions); err != nil {
		return err
	}

	for _, question := range questions {
		randNum := randomdata.Number(0, 4)
		for i := 0; i < 4; i++ {
			question_answer := models.QuestionAnswer{
				Content:   randomdata.Country(0),
				IsCorrect: i == randNum,
				Question:  question,
			}

			if err := tx.Create(&question_answer).Error; err != nil {
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
