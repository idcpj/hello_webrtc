package helpers

type Response struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	Data   interface{} `json:"data" `
}

func NewErrorResponse(msg string)*Response{
	return &Response{
		Status: 0,
		Msg:    msg,
		Data:   nil,
	}
}

func NewSuccessResponse(msg string,data interface{})*Response{
	return &Response{
		Status: 1,
		Msg:    msg,
		Data:   data,
	}
}
