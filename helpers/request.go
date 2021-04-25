package helpers

type Request struct {
	Type   string      `json:"type"`
	RoomId string      `json:"roomid"`
	Uid    string      `json:"uid"`
	Data   interface{} `json:"data"`
}
