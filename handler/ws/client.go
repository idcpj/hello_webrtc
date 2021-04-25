package ws

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func newClient(con *websocket.Conn,uid string,socket *socket)*client {
	return &client{
		con: con,
		uid: uid,
		socket:socket,
		quit: make(chan struct{},1),
	}
}


type IClient interface {
	read()
	write(msg []byte) error
	Close() error
	getId() string
}


type client struct {
	socket *socket
	con *websocket.Conn

	uid string
	quit chan struct{}
}


func (c *client) getId()string {
	return c.uid
}
func (c *client) read() {

	c.con.SetReadLimit(maxMessageSize)
	c.con.SetReadDeadline(time.Now().Add(pongWait))

	c.con.SetPongHandler(func(string) error {
		c.con.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {

		select {
		case <-c.quit:
			return
		default:

			msgType, message, err := c.con.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				c.Close()
				break
			}
			if msgType == websocket.TextMessage {
				message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
				// todo 根据类型发送给指定人员
				c.socket.conns.broadcast(message)
			}

		}

	}
}

func (c *client) write(msg []byte) error {
	c.con.SetWriteDeadline(time.Now().Add(writeWait))
	return c.con.WriteMessage(websocket.TextMessage,msg)
}

func (c *client) Close()error {
	log.Printf("client %s 退出 \n",c.getId())
	c.socket.conns.del(c)
	c.con.WriteMessage(websocket.CloseMessage, []byte{})
	close(c.quit)
	return  c.con.Close()
}

