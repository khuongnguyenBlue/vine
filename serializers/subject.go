package serializers

type Subject struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

type SubjectsList struct {
	Subjects []Subject `json:"subjects"`
}
