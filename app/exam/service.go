package exam

import (
	"github.com/khuongnguyenBlue/vine/dtos"
	"github.com/khuongnguyenBlue/vine/models"
	"net/http"
)

type Service interface {
	FetchBySubjectID(subjectID uint) ([]models.Exam, error)
	GetByID(id uint) (models.Exam, error)
	GetByIDWithQuestionsAnswers(id uint) (models.Exam, error)
	SaveSubmittedExam(submittedExam dtos.SubmittedExam, userID uint) (models.ExamResult, int)
	GetExamForReview(examID, userID uint) (dtos.ReviewExam, int)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) FetchBySubjectID(subjectID uint) ([]models.Exam, error) {
	return s.Repo.FetchBySubjectID(subjectID)
}

func (s *service) GetByID(id uint) (models.Exam, error) {
	return s.Repo.GetByID(id)
}

func (s *service) GetByIDWithQuestionsAnswers(id uint) (models.Exam, error) {
	return s.Repo.GetByIDWithQuestionsAnswers(id)
}

func (s *service) SaveSubmittedExam(submittedExam dtos.SubmittedExam, userID uint) (models.ExamResult, int) {
	exam, err := s.Repo.GetByIDWithQuestionsAnswers(submittedExam.ID)
	if err != nil {
		return models.ExamResult{}, http.StatusNotFound
	}

	var examResult = models.ExamResult{ExamID: submittedExam.ID, UserID: userID, SpentTime: submittedExam.SpentTime}
	for _, answeredQuestion := range submittedExam.AnsweredQuestions {
		targetQuestion, found := exam.GetQuestionByID(answeredQuestion.ID)
		if !found {
			return models.ExamResult{}, http.StatusBadRequest
		}

		correctAnswerID, found := targetQuestion.GetCorrectAnswerID()
		if !found {
			return models.ExamResult{}, http.StatusBadRequest
		}

		isCorrect := correctAnswerID == answeredQuestion.AnswerID
		if isCorrect {
			examResult.Score += 1
		}

		userAnswer := models.UserAnswer{
			QuestionID: answeredQuestion.ID,
			AnswerID:   answeredQuestion.AnswerID,
			IsCorrect:  isCorrect,
		}
		examResult.UserAnswers = append(examResult.UserAnswers, userAnswer)
	}

	examResult, err = s.Repo.CreateExamResult(examResult)
	if err != nil {
		return models.ExamResult{}, http.StatusInternalServerError
	}

	examResult.Exam = exam
	return examResult, 0
}

func (s *service) GetExamForReview(examID, userID uint) (dtos.ReviewExam, int) {
	exam, err := s.Repo.GetByIDWithQuestionsAnswers(examID)
	if err != nil {
		return dtos.ReviewExam{}, http.StatusNotFound
	}

	examResult, err := s.Repo.GetExamResult(examID, userID)
	if err != nil {
		return dtos.ReviewExam{}, http.StatusNotFound
	}

	var reviewExamDTO dtos.ReviewExam
	reviewExamDTO.Extract(exam, examResult)
	return reviewExamDTO, 0
}
