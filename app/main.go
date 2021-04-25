package main

import (
	"log"
	"os"
	"os/signal"
	"theia/handler/ws"
	"theia/router"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}

func main() {

	ws.Init()

	g := router.Router()

	g.Run(":8083")

	s:=make(chan os.Signal,1)
	signal.Notify(s,os.Kill,os.Interrupt)
	<-s
}
