package models

type QuestionAnswer struct {
	ID         uint   `json:"id"`
	Content    string `json:"content" gorm:"type:varchar(100); not null"`
	IsCorrect  bool   `json:"is_correct" gorm:"default:false"`
	QuestionID uint   `json:"question_id" gorm:"not null;index"`
	Question   Question
}
