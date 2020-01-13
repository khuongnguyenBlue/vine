package exam

import (
	"github.com/khuongnguyenBlue/vine/app/exam_result"
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
	GetExamRanking(examID uint) (dtos.ExamRanking, error)
}

type service struct {
	examRepo Repository
	examResultRepo exam_result.Repository
}

func NewService(examRepo Repository, examResultRepo exam_result.Repository) Service {
	return &service{examRepo: examRepo, examResultRepo: examResultRepo}
}

func (s *service) FetchBySubjectID(subjectID uint) ([]models.Exam, error) {
	return s.examRepo.FetchBySubjectID(subjectID)
}

func (s *service) GetByID(id uint) (models.Exam, error) {
	return s.examRepo.GetByID(id)
}

func (s *service) GetByIDWithQuestionsAnswers(id uint) (models.Exam, error) {
	return s.examRepo.GetByIDWithQuestionsAnswers(id)
}

func (s *service) SaveSubmittedExam(submittedExam dtos.SubmittedExam, userID uint) (models.ExamResult, int) {
	exam, err := s.examRepo.GetByIDWithQuestionsAnswers(submittedExam.ID)
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

	examResult, err = s.examResultRepo.Create(examResult)
	if err != nil {
		return models.ExamResult{}, http.StatusInternalServerError
	}

	examResult.Exam = exam
	return examResult, 0
}

func (s *service) GetExamForReview(examID, userID uint) (dtos.ReviewExam, int) {
	exam, err := s.examRepo.GetByIDWithQuestionsAnswers(examID)
	if err != nil {
		return dtos.ReviewExam{}, http.StatusNotFound
	}

	examResult, err := s.examResultRepo.GetByExamIDAndUserID(examID, userID)
	if err != nil {
		return dtos.ReviewExam{}, http.StatusNotFound
	}

	var reviewExamDTO dtos.ReviewExam
	reviewExamDTO.Extract(exam, examResult)
	return reviewExamDTO, 0
}

func (s *service) GetExamRanking(examID uint) (dtos.ExamRanking, error) {
	exam, err := s.examRepo.GetByIDWithExamResults(examID)
	if err != nil {
		return dtos.ExamRanking{}, err
	}

	var examRankingDTO dtos.ExamRanking
	examRankingDTO.Extract(exam)
	return examRankingDTO, nil
}
