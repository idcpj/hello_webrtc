package ws

import (
	"errors"
	"sync"
)


var (
	ERRPR_USER_NOT_EXIST=errors.New("用户不存在")
)

type IHub interface {
	add(client IClient) error
	del(uid string) error
	send(uid string,msg []byte) error
	exist(uid string) bool
	broadcast(msg []byte) error
	Close() error
}



var (
	ERROR_USER_EXIST = errors.New("userid 已经存在")
)

func newHub() *hub {
	return &hub{
		conns:make(map[string]IClient),
	}
}

type hub struct {
	conns map[string]IClient
	mx    sync.RWMutex
}

func (h *hub) del(uid string) error {
	if h.exist(uid) {
		delete(h.conns,uid)
	}
	return nil
}

func (h *hub) send(uid string, msg []byte) error {
	if !h.exist(uid) {
		return ERRPR_USER_NOT_EXIST
	}

	return h.conns[uid].write(msg)
}

func (h *hub) exist(uid string) bool {
	_,ok:=h.conns[uid]
	return ok
}

func (h *hub) Del(client IClient) error {
	h.mx.Lock()
	defer h.mx.Unlock()

	delete(h.conns,client.getId())
	return nil
}


func (h *hub) add(client IClient) error {
	h.mx.Lock()
	defer h.mx.Unlock()

	_, ok := h.conns[client.getId()]
	if ok {
		return ERROR_USER_EXIST
	}

	h.conns[client.getId()] = client
	return nil
}


func (h *hub) broadcast(message []byte) error {
	h.mx.RLock()
	defer h.mx.RUnlock()

	for _, client := range h.conns {
		client.write(message)
	}

	return nil
}


func (h *hub) Close() error  {
	h.mx.Lock()
	defer h.mx.Unlock()

	for _, con := range h.conns {
		con.Close()
	}

	return nil
}