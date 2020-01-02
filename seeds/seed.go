package seeds

import (
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/models"
	"log"
)

func All(c chan error) {
	if ssErr := seedSubjects(); ssErr != nil {
		c <- ssErr
		return
	}

	// other seeding ...

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
