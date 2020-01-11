package exam

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

func (r *repository) FetchBySubjectID(subjectId uint) ([]models.Exam, error)  {
	var exams []models.Exam
	err := r.Conn.Preload("Questions").Where("subject_id = ?", subjectId).Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

func (r *repository) GetByID(id uint) (models.Exam, error) {
	var exam = models.Exam{ID: id}
	if err := r.Conn.Preload("Questions").First(&exam).Error; err != nil {
		return models.Exam{}, err
	}

	return exam, nil
}

func (r *repository) GetByIDWithQuestionsAnswers(id uint) (models.Exam, error) {
	var exam = models.Exam{ID: id}
	if err := r.Conn.Preload("Questions").Preload("Questions.QuestionAnswers").
		First(&exam).Error; err != nil {
		return models.Exam{}, err
	}

	return exam, nil
}

func (r *repository) CreateExamResult(examResult models.ExamResult) (models.ExamResult, error) {
	if err := r.Conn.Create(&examResult).Error; err != nil {
		return models.ExamResult{}, err
	}

	return examResult, nil
}
