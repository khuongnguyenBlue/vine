package models

type UserAnswer struct {
	ID uint
	ExamResultID uint `gorm:"not null;index"`
	ExamResult ExamResult
	QuestionID uint `gorm:"not null;index"`
	Question Question
	AnswerID uint `gorm:"not null;index"`
	Answer QuestionAnswer
	IsCorrect bool `gorm:"default:false"`
}
