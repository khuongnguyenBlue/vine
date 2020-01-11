package models

type Question struct {
	ID        uint
	Content   string `gorm:"type:varchar(500); not null"`
	SubjectID uint   `gorm:"not null;index"`
	Subject   Subject
	QuestionAnswers []QuestionAnswer
}

func (q *Question) GetCorrectAnswerID() (uint, bool) {
	for _, answer := range q.QuestionAnswers {
		if answer.IsCorrect == true {
			return answer.ID, true
		}
	}

	return 0, false
}
