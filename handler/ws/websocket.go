package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"theia/helpers"
)

var (
	_socket *socket
)

func Init() {

	log.Printf("websocket 初始化")

	_socket = newSocket()
}

func newSocket() *socket {

	h := newHub()
	return &socket{
		s: &websocket.Upgrader{
			HandshakeTimeout:  30,
			ReadBufferSize:    1024,
			WriteBufferSize:   1024,
			EnableCompression: true,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		conns: h,
	}
}

type socket struct {
	s     *websocket.Upgrader
	conns IHub
}

func (s *socket) newConn(w http.ResponseWriter, r *http.Request) error {
	con, err := s.s.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	query := r.URL.Query()
	uid := query.Get("uid")
	if uid == "" {
		return ERROR_UID_NOT_EXIST
	}

	client := newClient(con, uid, s)

	if err = s.conns.add(client); err != nil {
		client.Write(helpers.NewErrorResp("", uid, SOCKET_LOGIN, err.Error()))
		return ERROR_USER_EXIST
	}

	go client.Run()

	return nil
}

func (s *socket) Close() error {
	return s.conns.Close()
}
