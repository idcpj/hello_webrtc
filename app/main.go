package main

import (
	"log"
	"os"
	"os/signal"
	"theia/handler/ws"
	"theia/router"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	ws.Init()

	g := router.Router()

	err := g.RunTLS(":8083", "ssl/server.crt", "ssl/server.key")
	//err := g.Run(":8083")
	if err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Kill, os.Interrupt)
	<-s
}
