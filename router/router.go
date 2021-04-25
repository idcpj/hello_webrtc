package router

import (
	"github.com/gin-gonic/gin"
	"theia/handler"
	"theia/handler/ws"
	"theia/middlewares"
)

func Router() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(middlewares.Allow())
	r.Use(middlewares.Auth())

	r.Static("/web", "./web")
	r.GET("/", handler.DefaultHomePageHandler)

	r.GET("/ws", ws.WebSocketServer)
	r.GET("/wss", ws.WebSocketServer)

	return r
}
