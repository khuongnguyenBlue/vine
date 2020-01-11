package dtos

import "github.com/khuongnguyenBlue/vine/models"

type Exam struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	TimeAllow      uint              `json:"time_allow"`
	Status         models.ExamStatus `json:"status"`
	QuestionsCount int               `json:"questions_count"`
	SubjectID      uint              `json:"subject_id"`
}

func (e *Exam) Extract(exam models.Exam) {
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

func (el *ExamsList) Extract(exams []models.Exam) {
	for _, exam := range exams {
		var e Exam
		e.Extract(exam)
		el.Exams = append(el.Exams, e)
	}
}

type FullExam struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	TimeAllow uint              `json:"time_allow"`
	Status    models.ExamStatus `json:"status"`
	SubjectID uint              `json:"subject_id"`
	Questions QuestionsList     `json:"questions"`
}

func (fe *FullExam) Extract(exam models.Exam) {
	fe.ID = exam.ID
	fe.Name = exam.Name
	fe.TimeAllow = exam.TimeAllow
	fe.Status = exam.Status
	fe.SubjectID = exam.SubjectID
	var questionsList QuestionsList
	questionsList.Extract(exam.Questions)
	fe.Questions = questionsList
}

type SubmittedExam struct {
	ID                uint                 `json:"id"`
	SpentTime uint `json:"spent_time" binding:"required"`
	AnsweredQuestions []QuestionWithAnswer `json:"answered_questions" binding:"required"`
}

type BriefResult struct {
	Score uint `json:"score"`
	SpentTime uint `json:"spent_time"`
}

func (er *BriefResult) Extract(result models.ExamResult) {
	er.Score = result.Score
	er.SpentTime = result.SpentTime
}
