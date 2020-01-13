package dtos

import "github.com/khuongnguyenBlue/vine/models"

type BaseExam struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	TimeAllow      uint              `json:"time_allow"`
	Status         models.ExamStatus `json:"status"`
	SubjectID      uint              `json:"subject_id"`
}

func (be *BaseExam) Extract(exam models.Exam) {
	be.ID = exam.ID
	be.Name = exam.Name
	be.TimeAllow = exam.TimeAllow
	be.Status = exam.Status
	be.SubjectID = exam.SubjectID
}

type Exam struct {
	BaseExam
	QuestionsCount int               `json:"questions_count"`
}

func (e *Exam) Extract(exam models.Exam) {
	var be BaseExam
	be.Extract(exam)
	e.BaseExam = be
	e.QuestionsCount = len(exam.Questions)
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
	BaseExam
	Questions QuestionsList     `json:"questions"`
}

func (fe *FullExam) Extract(exam models.Exam) {
	var be BaseExam
	be.Extract(exam)
	fe.BaseExam = be
	var questionsList QuestionsList
	questionsList.Extract(exam.Questions)
	fe.Questions = questionsList
}

type SubmittedExam struct {
	ID                uint                 `json:"id"`
	SpentTime         uint                 `json:"spent_time" binding:"required"`
	AnsweredQuestions []QuestionWithAnswer `json:"answered_questions" binding:"required"`
}

type BriefResult struct {
	Score     uint `json:"score"`
	SpentTime uint `json:"spent_time"`
}

func (er *BriefResult) Extract(result models.ExamResult) {
	er.Score = result.Score
	er.SpentTime = result.SpentTime
}

type ReviewExam struct {
	BaseExam
	BriefResult
	ReviewQuestions []ReviewQuestion  `json:"review_questions"`
}

func (re *ReviewExam) Extract(exam models.Exam, examResult models.ExamResult)  {
	var be BaseExam
	be.Extract(exam)
	re.BaseExam = be

	var br BriefResult
	br.Extract(examResult)
	re.BriefResult = br

	for _, question := range exam.Questions {
		userAnswerID := examResult.GetUserAnswerID(question.ID)
		var reviewQuestion ReviewQuestion
		reviewQuestion.Extract(question, userAnswerID)
		re.ReviewQuestions = append(re.ReviewQuestions, reviewQuestion)
	}
}

type UserRecord struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Score uint `json:"score"`
	SpentTime uint `json:"spent_time"`
}

func (ur *UserRecord) Extract(result models.ExamResult)  {
	ur.ID = result.ID
	ur.UserID = result.UserID
	ur.Score = result.Score
	ur.SpentTime = result.SpentTime
}

type ExamRanking struct {
	BaseExam
	UserRecords []UserRecord `json:"user_records"`
}

func (er *ExamRanking) Extract(exam models.Exam) {
	var be BaseExam
	be.Extract(exam)
	er.BaseExam = be

	for _, result := range exam.ExamResults {
		var ur UserRecord
		ur.Extract(result)
		er.UserRecords = append(er.UserRecords, ur)
	}
}
