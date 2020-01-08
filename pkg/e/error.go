package e

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func NewError(errCode int) ErrorResponse {
	return ErrorResponse{Code: errCode, Message: GetMsg(errCode)}
}
