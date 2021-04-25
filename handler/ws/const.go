package ws

import (
	"errors"
)

var (
	ERRPR_USER_NOT_EXIST=errors.New("用户不存在")
	ERROR_USER_EXIST = errors.New("userid 已经存在")
	ERROR_UID_NOT_EXIST =errors.New("userid 不存在")
	ERROR_REQUEST_NOT_ALLOW =errors.New("该请求不是 websocket 连接")
)