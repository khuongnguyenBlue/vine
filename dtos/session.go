package dtos

type LoginAccount struct {
	PhoneNumber string `binding:"required,lte=11" json:"phone_number" form:"phone_number"`
	Password    string `binding:"required,lte=20" json:"password" form:"password"`
}

type RegisterInfo struct {
	PhoneNumber     string `binding:"required,lte=11" json:"phone_number" form:"phone_number"`
	Password        string `binding:"required,lte=20" json:"password" form:"password"`
	ConfirmPassword string `binding:"required,lte=20,eqfield=Password" json:"confirm_password" form:"confirm_password"`
}
