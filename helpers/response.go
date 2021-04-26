package helpers

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`

	Type   string      `json:"type"`
	RoomId string      `json:"roomid"`
	Uid    string      `json:"uid"`
	Data   interface{} `json:"data" `
}

func NewReqToResp(request *Request) *Response {
	return &Response{
		Status: 1,
		RoomId: request.RoomId,
		Uid:    request.Uid,
		Type:   request.Type,
		Data:   request.Data,
	}
}

func NewErrorResp(roomid string, uid string, Type string, msg string) *Response {
	return &Response{
		Status: 0,
		Msg:    msg,
		RoomId: roomid,
		Uid:    uid,
		Type:   Type,
		Data:   nil,
	}
}

func NewSuccessResp(roomid string, uid string, Type string, data interface{}) *Response {
	return &Response{
		Status: 1,
		RoomId: roomid,
		Uid:    uid,
		Type:   Type,
		Data:   data,
	}
}
