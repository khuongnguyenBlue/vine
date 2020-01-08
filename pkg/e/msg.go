package e

var MsgFlags = map[int]string{
	InvalidParams:          "Tham số không hợp lệ",
	NotFound:               "Không tìm thấy tài nguyên",
	InternalServerError:    "Có lỗi xảy ra",
	ExistedPhoneNumber:     "Số điện thoại này đã được sử dụng để đăng ký",
	InvalidRegisterAccount: "Thông tin đăng ký không hợp lệ",
	InvalidLoginAccount:    "Thông tin đăng nhập không hợp lệ",
	WrongLoginAccount:      "Thông tin đăng nhập không chính xác",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InternalServerError]
}
