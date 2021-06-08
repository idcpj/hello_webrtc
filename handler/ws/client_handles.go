package ws

import (
	"theia/helpers"
)

func (c *client) baseSendOther(request *helpers.Request) {
	c.socket.conns.sendOther(request.RoomId, helpers.NewReqToResp(request))
}

func (c *client) baseCallBack(request *helpers.Request) {
	c.SuccessResp(request, request.Data)
}

func (c *client) errorCallBack(request *helpers.Request) {
	c.SuccessResp(request, request.Data)
}

func (c *client) roomJoin(request *helpers.Request) {
	if request.RoomId == "" {
		c.ErrorResp(request, ERROR_ROOMID_NOT_EXIST)
		return
	}
	err := c.socket.conns.join(request.RoomId, c.uid)
	if err != nil {
		c.ErrorResp(request, err)
		return
	}
	c.SuccessResp(request, request.Data)
}

func (c *client) quitRoom(request *helpers.Request) {
	err := c.socket.conns.quit(request.RoomId, request.Uid)
	if err != nil {
		c.ErrorResp(request, err)
		return
	}
	c.SuccessResp(request, nil)
}
