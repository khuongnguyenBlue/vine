package dtos

import "github.com/khuongnguyenBlue/vine/models"

type QuestionAnswer struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func (qa *QuestionAnswer) Extract(answer models.QuestionAnswer) {
	qa.ID = answer.ID
	qa.Content = answer.Content
}

type AnswersList []QuestionAnswer

func (al *AnswersList) Extract(answers []models.QuestionAnswer) {
	for _, answer := range answers {
		var qa QuestionAnswer
		qa.Extract(answer)
		*al = append(*al, qa)
	}
}

type ReviewAnswer struct {
	QuestionAnswer
	IsCorrect bool `json:"is_correct"`
	IsChosen  bool `json:"is_chosen"`
}

func (ra *ReviewAnswer) Extract(answer models.QuestionAnswer, userAnswerID uint) {
	var qa QuestionAnswer
	qa.Extract(answer)
	ra.QuestionAnswer = qa
	ra.IsCorrect = answer.IsCorrect
	ra.IsChosen = answer.ID == userAnswerID
}
