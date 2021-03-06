package dtos

import "github.com/khuongnguyenBlue/vine/models"

type Question struct {
	ID              uint        `json:"id"`
	Content         string      `json:"content"`
	QuestionAnswers AnswersList `json:"question_answers"`
}

func (q *Question) Extract(question models.Question) {
	q.ID = question.ID
	q.Content = question.Content
	var answersList AnswersList
	answersList.Extract(question.QuestionAnswers)
	q.QuestionAnswers = answersList
}

type QuestionsList []Question

func (ql *QuestionsList) Extract(questions []models.Question) {
	for _, question := range questions {
		var q Question
		q.Extract(question)
		*ql = append(*ql, q)
	}
}

type QuestionWithAnswer struct {
	ID       uint `json:"id"`
	AnswerID uint `json:"answer_id"`
}

type ReviewQuestion struct {
	ID            uint           `json:"id"`
	Content       string         `json:"content"`
	ReviewAnswers []ReviewAnswer `json:"review_answers"`
}

func (rq *ReviewQuestion) Extract(question models.Question, userAnswerID uint) {
	rq.ID = question.ID
	rq.Content = question.Content

	for _, answer := range question.QuestionAnswers {
		var reviewAnswer ReviewAnswer
		reviewAnswer.Extract(answer, userAnswerID)
		rq.ReviewAnswers = append(rq.ReviewAnswers, reviewAnswer)
	}
}
