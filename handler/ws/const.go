package ws

import (
	"errors"
)

var (
	// user
	ERRPR_USER_NOT_EXIST = errors.New("用户不存在")
	ERROR_USER_EXIST     = errors.New("userid 已经存在")
	ERROR_UID_NOT_EXIST  = errors.New("userid 不存在")

	// room

	ERROR_ROOM_NOT_EXIST = errors.New("房间不存在")

	ERROR_REQUEST_NOT_ALLOW = errors.New("该请求不是 websocket 连接")
)

const (
	// peer
	PEER_CANDIDATE = "candidate"
	PEER_ANSWER    = "answer"
	PEER_OFFER     = "offer"

	PEER_READY = "peer_ready"

	// room
	ROOM_JOIN = "room_join"
	ROOM_QUIT = "room_quit"

	// message
	SEND_MSG = "send_msg"

	SOCKET_HEART = "heart"
	SOCKET_LOGIN = "login"
)
