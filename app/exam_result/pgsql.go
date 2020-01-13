package exam_result

import (
	"github.com/jinzhu/gorm"
	"github.com/khuongnguyenBlue/vine/models"
)

type repository struct {
	Conn *gorm.DB
}

func NewPgsqlRepository(conn *gorm.DB) Repository {
	return &repository{Conn: conn}
}

func (r *repository) Create(examResult models.ExamResult) (models.ExamResult, error) {
	if err := r.Conn.Create(&examResult).Error; err != nil {
		return models.ExamResult{}, err
	}

	return examResult, nil
}

func (r *repository) GetByExamIDAndUserID(examID, userID uint) (models.ExamResult, error)  {
	var examResult models.ExamResult
	if err := r.Conn.Preload("UserAnswers").Where("exam_id = ? AND user_id = ?", examID, userID).
		First(&examResult).Error; err != nil {
		return models.ExamResult{}, err
	}

	return examResult, nil
}
