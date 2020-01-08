package models

type Subject struct {
	ID        uint
	Name      string `gorm:"type:varchar(20); unique; not null"`
	Questions []Question
}
