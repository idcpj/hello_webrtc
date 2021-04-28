package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"theia/helpers"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
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
	Run()

	Read()
	Write(response *helpers.Response) error

	ErrorResp(response *helpers.Request, err error)
	SuccessResp(response *helpers.Request, data interface{})

	match(request *helpers.Request)
	InitHandles()

	GetUid() string
	SetRoomId(string)

	Close() error
}

type client struct {
	socket *socket
	con    *websocket.Conn
	onece  sync.Once

	uid    string
	roomid string
	quit   chan struct{}
	mux    map[string]func(request *helpers.Request)
}

func (c *client) SetRoomId(roomid string) {
	c.roomid = roomid
}

func (c *client) Run() {
	c.InitHandles()
	go c.Read()
	c.Heart()
}
func (c *client) GetUid() string {
	return c.uid
}

func (c *client) Read() {
	defer func() {
		e := recover()
		if e != nil {
			log.Println(e)
		}

		c.Close()
	}()

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
			var request = &helpers.Request{}
			err := c.con.ReadJSON(request)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				} else {
					log.Printf("error: %v", err)
				}
				return
			}

			c.match(request)
		}

	}
}

func (c *client) match(request *helpers.Request) {
	call, ok := c.mux[request.Type]
	if !ok {
		log.Println("指令类型不存在")
	}
	call(request)
}

func (c *client) InitHandles() {
	c.mux = make(map[string]func(request *helpers.Request))
	c.mux[PEER_ANSWER] = c.baseSendOther
	c.mux[PEER_CANDIDATE] = c.baseSendOther
	c.mux[PEER_OFFER] = c.baseSendOther
	c.mux[PEER_READY] = c.baseCallBack

	c.mux[ROOM_JOIN] = c.roomJoin
	c.mux[ROOM_QUIT] = c.quitRoom

	c.mux[SOCKET_HEART] = c.baseCallBack
}

func (c *client) Write(response *helpers.Response) error {
	_ = c.con.SetWriteDeadline(time.Now().Add(writeWait))
	e := c.con.WriteJSON(response)
	if e != nil {
		log.Println(e)

		c.Close()
	}
	return nil
}

func (c *client) ErrorResp(request *helpers.Request, err error) {
	c.Write(helpers.NewErrorResp(request.RoomId, request.Uid, request.Type, err.Error()))
}

func (c *client) SuccessResp(request *helpers.Request, data interface{}) {
	c.Write(helpers.NewSuccessResp(request.RoomId, request.Uid, request.Type, data))
}

func (c *client) Close() error {
	c.onece.Do(func() {
		log.Printf("client %s 退出 \n", c.GetUid())

		if c.roomid != "" {
			c.socket.conns.quit(c.roomid, c.GetUid())
		}
		c.socket.conns.del(c.GetUid())
		_ = c.con.WriteMessage(websocket.CloseMessage, []byte{})
		close(c.quit)
		c.con.Close()
	})
	return nil

}

func (c *client) Heart() {
	tick := time.Tick(pingPeriod)
	for {
		select {
		case <-c.quit:
			return
		case <-tick:
			e := c.con.WriteMessage(websocket.PingMessage, nil)
			if e != nil {
				log.Println(e)
				c.Close()
			}
		}
	}
}
