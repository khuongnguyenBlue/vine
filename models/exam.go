package models

type ExamStatus int

const (
	Waiting  ExamStatus = 0
	Opening  ExamStatus = 1
	Finished ExamStatus = 2
)

type Exam struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name" gorm:"type:varchar(50); not null"`
	TimeAllow uint       `json:"time_allow" gorm:"not null"`
	Status    uint       `json:"status" gorm:"default:0"`
	Questions []Question `gorm:"many2many:exam_questions"`
	SubjectID uint       `json:"subject_id" gorm:"not null"`
	Subject   Subject
}
