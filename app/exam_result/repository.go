package exam_result

import "github.com/khuongnguyenBlue/vine/models"

type Repository interface {
	Create(examResult models.ExamResult) (models.ExamResult, error)
	GetByExamIDAndUserID(examID, userID uint) (models.ExamResult, error)
}
