package dtos

import "github.com/khuongnguyenBlue/vine/models"

type Exam struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	TimeAllow      uint   `json:"time_allow"`
	Status         uint   `json:"status"`
	QuestionsCount int    `json:"questions_count"`
	SubjectID      uint   `json:"subject_id"`
}

func (e *Exam) Extract(exam models.Exam)  {
	e.ID = exam.ID
	e.Name = exam.Name
	e.TimeAllow = exam.TimeAllow
	e.Status = exam.Status
	e.QuestionsCount = len(exam.Questions)
	e.SubjectID = exam.SubjectID
}

type ExamsList struct {
	Exams []Exam `json:"exams"`
}

func (el *ExamsList) Extract(exams []models.Exam)  {
	for _, exam := range exams {
		var e Exam
		e.Extract(exam)
		el.Exams = append(el.Exams, e)
	}
}

