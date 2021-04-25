package router

import (
	"github.com/gin-gonic/gin"
	"theia/handler"
	"theia/handler/ws"
)

func Router() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r:=gin.Default()

	r.Static("/web/","./web")
	r.GET("/", handler.DefaultHomePageHandler)



	r.GET("/ws", ws.WebSocketServer)
	r.GET("/wss", ws.WebSocketServer)


	return r
}
