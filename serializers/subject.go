package serializers

import "github.com/khuongnguyenBlue/vine/models"

type SubjectJson struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *SubjectJson) Parse(subject models.Subject) {
	s.ID = subject.ID
	s.Name = subject.Name
}

type SubjectsJson struct {
	Subjects []SubjectJson `json:"subjects"`
}

func (sl *SubjectsJson) Parse(subjects []models.Subject)  {
	for _, subject := range subjects {
		var s SubjectJson;
		s.Parse(subject)
		sl.Subjects = append(sl.Subjects, s)
	}
}
