package ws

import (
	"github.com/gin-gonic/gin"
	"log"
)

func WebSocketServer(c *gin.Context) {
	if !c.IsWebsocket() {
		log.Println(ERROR_REQUEST_NOT_ALLOW)
	}

	if err := _socket.newConn(c.Writer, c.Request); err != nil {
		log.Println(err)
	}
	return
}
