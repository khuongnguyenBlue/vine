package models

type Question struct {
	ID        uint
	Content   string `gorm:"type:varchar(500); not null"`
	SubjectID uint   `gorm:"not null"`
	Subject   Subject
	QuestionAnswers []QuestionAnswer
}
