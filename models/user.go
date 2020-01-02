package models

type User struct {
	ID                uint
	Name              string `gorm:"type:varchar(50)"`
	PhoneNumber       string `gorm:"type:varchar(11); unique_index; not null"`
	EncryptedPassword string
}
