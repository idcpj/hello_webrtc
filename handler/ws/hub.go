package ws

import (
	"log"
	"sync"
	"theia/helpers"
)

//IHub
type IHub interface {

	// user
	add(client IClient) error
	del(uid string) error
	exist(uid string) bool

	//room
	join(roomId string, uid string) error
	quit(roomId string, uid string) error
	roomExist(roomId string) bool
	roomEmpty(roomId string) bool
	roomIsMax(roomid string) bool
	MemberIsExist(room string, uid string) bool

	// msg
	send(uid string, response *helpers.Response) error
	broadcast(roomId string, response *helpers.Response) error
	sendOther(roomId string, response *helpers.Response) error

	Close() error
}

func newHub() *hub {
	return &hub{
		conns:    make(map[string]IClient),
		rooms:    make(map[string]map[string]IClient),
		roomsMax: 2,
	}
}

type hub struct {
	conns map[string]IClient
	rooms map[string]map[string]IClient

	roomsMax int
	mx       sync.RWMutex
}

func (h *hub) MemberIsExist(room string, uid string) bool {
	h.mx.RLock()
	defer h.mx.RUnlock()
	if len(h.rooms[room]) <= 0 {
		return false
	}
	_, ok := h.rooms[room][uid]
	if ok {
		return true
	}
	return false
}

func (h *hub) roomIsMax(roomid string) bool {
	h.mx.RLock()
	defer h.mx.RUnlock()

	if len(h.rooms[roomid]) >= h.roomsMax {
		return true
	}
	return false
}

func (h *hub) join(roomId string, uid string) error {

	if h.roomEmpty(roomId) {
		h.rooms[roomId] = make(map[string]IClient)
	} else if h.MemberIsExist(roomId, uid) {
		return ERROR_ROOM_MEMBER_IS_EXIST
	} else if h.roomIsMax(roomId) {
		return ERROR_ROOM_MEMBER_TOO_MANY
	}

	h.mx.Lock()
	defer h.mx.Unlock()

	client, ok := h.conns[uid]
	if !ok {
		return ERROR_UID_NOT_EXIST
	}

	h.rooms[roomId][uid] = client
	client.SetRoomId(roomId)

	return nil

}

func (h *hub) quit(roomId string, uid string) error {
	if h.roomEmpty(roomId) {
		return ERROR_ROOM_NOT_EXIST
	}

	h.mx.Lock()
	defer h.mx.Unlock()

	delete(h.rooms[roomId], uid)
	return nil

}

func (h *hub) roomExist(roomId string) bool {
	h.mx.RLock()
	defer h.mx.RUnlock()
	_, ok := h.rooms[roomId]
	return ok

}

func (h *hub) roomEmpty(roomId string) bool {
	h.mx.RLock()
	defer h.mx.RUnlock()
	return len(h.rooms[roomId]) == 0
}

func (h *hub) del(uid string) error {
	if h.exist(uid) {
		delete(h.conns, uid)
	}
	return nil
}

func (h *hub) send(uid string, response *helpers.Response) error {
	if !h.exist(uid) {
		return ERRPR_USER_NOT_EXIST
	}

	return h.conns[uid].Write(response)
}

func (h *hub) exist(uid string) bool {
	_, ok := h.conns[uid]
	return ok
}

func (h *hub) Del(client IClient) error {
	h.mx.Lock()
	defer h.mx.Unlock()

	delete(h.conns, client.GetUid())
	return nil
}

func (h *hub) add(client IClient) error {
	h.mx.Lock()
	defer h.mx.Unlock()

	_, ok := h.conns[client.GetUid()]
	if ok {
		return ERROR_USER_EXIST
	}

	h.conns[client.GetUid()] = client
	log.Println("添加人员后,总人数为", len(h.conns))

	return nil
}

func (h *hub) sendOther(roomId string, response *helpers.Response) error {
	h.mx.RLock()
	defer h.mx.RUnlock()

	for _, client := range h.rooms[roomId] {
		if client.GetUid() != response.Uid {
			client.Write(response)
		}
	}

	return nil
}
func (h *hub) broadcast(roomId string, response *helpers.Response) error {
	h.mx.RLock()
	defer h.mx.RUnlock()

	for _, client := range h.rooms[roomId] {
		client.Write(response)
	}

	return nil
}

func (h *hub) Close() error {
	h.mx.Lock()
	defer h.mx.Unlock()

	for _, con := range h.conns {
		con.Close()
	}

	return nil
}
