package models

type ExamResult struct {
	ID          uint
	ExamID      uint `gorm:"not null;index"`
	Exam        Exam
	UserID      uint `gorm:"not null;index"`
	User        User
	Score       uint `gorm:"default:0"`
	SpentTime   uint `gorm:"not null;index"`
	UserAnswers []UserAnswer
}

func (er *ExamResult) GetUserAnswerID(questionID uint) uint {
	for _, ua := range er.UserAnswers {
		if ua.QuestionID == questionID {
			return ua.AnswerID
		}
	}

	return 0
}
