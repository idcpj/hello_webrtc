package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"theia/helpers"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 60 * time.Second

	// Time allowed to ReadJson the next pong message from the peer.
	pongWait = 60 * time.Second

	// send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)

func newClient(con *websocket.Conn, uid string, socket *socket) *client {
	return &client{
		con:    con,
		uid:    uid,
		socket: socket,
		quit:   make(chan struct{}, 1),
	}
}

type IClient interface {
	ReadJson()
	WriteJSON(response helpers.Response) error
	Close() error
	getId() string
}

type client struct {
	socket *socket
	con    *websocket.Conn

	uid  string
	quit chan struct{}
}

func (c *client) getId() string {
	return c.uid
}

func (c *client) ReadJson() {

	//c.con.SetReadLimit(maxMessageSize)
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
			var request = helpers.Request{}
			err := c.con.ReadJSON(&request)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				} else {
					log.Printf("error: %v", err)
				}

				c.Close()
				break
			}

			c.handle(request)
		}

	}
}

func (c *client) handle(request helpers.Request) {
	switch request.Type {
	case PEER_ANSWER:
		fmt.Printf("%+v\n", request.Type)
		c.socket.conns.broadcast(request.RoomId, helpers.NewReqToResp(request))
	case PEER_CANDIDATE:
		c.socket.conns.broadcast(request.RoomId, helpers.NewReqToResp(request))
	case PEER_OFFER:
		fmt.Printf("%+v\n", request.Type)
		c.socket.conns.broadcast(request.RoomId, helpers.NewReqToResp(request))

	case PEER_READY:
		c.SuccessResponse(request, request.Data)

	case ROOM_JOIN:
		err := c.socket.conns.join(request.RoomId, request.Uid)
		if err != nil {
			c.ErrorResponse(request, err)
			return
		}
		c.SuccessResponse(request, request.Data)

	case ROOM_QUIT:
		err := c.socket.conns.quit(request.RoomId, request.Uid)
		if err != nil {
			c.ErrorResponse(request, err)
			return
		}
		c.SuccessResponse(request, nil)

	default:
		log.Println("位置指令")

	}
}

func (c *client) WriteJSON(response helpers.Response) error {
	_ = c.con.SetWriteDeadline(time.Now().Add(writeWait))
	e := c.con.WriteJSON(response)
	if e != nil {
		log.Println(e)
	}
	return nil
}

func (c *client) ErrorResponse(request helpers.Request, err error) {
	c.WriteJSON(helpers.NewErrorResp(request.RoomId, request.Uid, request.Type, err.Error()))
}

func (c *client) SuccessResponse(request helpers.Request, data interface{}) {
	c.WriteJSON(helpers.NewSuccessResp(request.RoomId, request.Uid, request.Type, data))
}

func (c *client) Close() error {
	log.Printf("client %s 退出 \n", c.getId())
	c.socket.conns.del(c.getId())
	_ = c.con.WriteMessage(websocket.CloseMessage, []byte{})
	close(c.quit)
	return c.con.Close()
}
