package models

type User struct {
	ID                uint   `json:"id"`
	Name              string `json:"name" gorm:"type:varchar(50)"`
	PhoneNumber       string `json:"phone_number" gorm:"type:varchar(11); unique_index; not null"`
	EncryptedPassword string `json:"encrypted_password" gorm:"type:varchar(11)"`
}
