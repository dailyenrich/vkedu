package response

type Response struct {
	Code uint `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Resp(code uint, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
