package ws

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"theia/helpers"
)



func WebSocketServer(c *gin.Context){
	if !c.IsWebsocket() {
		c.JSON(http.StatusNotAcceptable,helpers.NewErrorResponse("该请求不是 websocket 连接"))
	}

	if err := _socket.newConn(c.Writer, c.Request);err!=nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest,helpers.NewErrorResponse(err.Error()))
	}
}
