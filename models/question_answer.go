package models

type QuestionAnswer struct {
	ID         uint
	Content    string `gorm:"type:varchar(100); not null"`
	IsCorrect  bool   `gorm:"default:false"`
	QuestionID uint   `gorm:"not null;index"`
	Question   Question
}
