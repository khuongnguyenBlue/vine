package serializers

type Exam struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	TimeAllow      uint   `json:"time_allow"`
	Status         uint   `json:"status"`
	QuestionsCount int    `json:"questions_count"`
	SubjectID      uint   `json:"subject_id"`
}

type ExamsList struct {
	Exams []Exam `json:"exams"`
}
