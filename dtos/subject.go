package dtos

import "github.com/khuongnguyenBlue/vine/models"

type Subject struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *Subject) Extract(subject models.Subject) {
	s.ID = subject.ID
	s.Name = subject.Name
}

type SubjectsList struct {
	Subjects []Subject `json:"subjects"`
}

func (sl *SubjectsList) Extract(subjects []models.Subject)  {
	for _, subject := range subjects {
		var s Subject;
		s.Extract(subject)
		sl.Subjects = append(sl.Subjects, s)
	}
}
