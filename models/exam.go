package models

type ExamStatus uint

const (
	Waiting  ExamStatus = 0
	Opening  ExamStatus = 1
	Finished ExamStatus = 2
)

type Exam struct {
	ID        uint
	Name      string     `gorm:"type:varchar(50); not null"`
	TimeAllow uint       `gorm:"not null"`
	Status    ExamStatus       `gorm:"default:0"`
	Questions []Question `gorm:"many2many:exam_questions"`
	SubjectID uint       `gorm:"not null"`
	Subject   Subject
}
