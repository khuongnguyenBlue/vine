package serializers

import "github.com/khuongnguyenBlue/vine/models"

type ExamJson struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	TimeAllow      uint   `json:"time_allow"`
	Status         uint   `json:"status"`
	QuestionsCount int    `json:"questions_count"`
	SubjectID      uint   `json:"subject_id"`
}

func (e *ExamJson) Parse(exam models.Exam)  {
	e.ID = exam.ID
	e.Name = exam.Name
	e.TimeAllow = exam.TimeAllow
	e.Status = exam.Status
	e.QuestionsCount = len(exam.Questions)
	e.SubjectID = exam.SubjectID
}

type ExamsJson struct {
	Exams []ExamJson `json:"exams"`
}

func (el *ExamsJson) Parse(exams []models.Exam)  {
	for _, exam := range exams {
		var e ExamJson
		e.Parse(exam)
		el.Exams = append(el.Exams, e)
	}
}

