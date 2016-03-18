package main

import (
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

type messaging struct {
	sockets map[int]*websocket.Conn
}

func (m *messaging) register(ws *websocket.Conn) int {
	if m.sockets == nil {
		m.sockets = make(map[int]*websocket.Conn)
	}

	id := rand.Int()
	fmt.Println("registering ", id)
	m.sockets[id] = ws
	return id
}

func (m *messaging) unregister(id int) {
	fmt.Println("unregistering ", id)
	delete(m.sockets, id)
}

func (m *messaging) sendMessage(mess message) error {
	ws := m.sockets[mess.ID]
	if ws == nil {
		return fmt.Errorf("id %d not found", mess.ID)
	}

	ws.WriteMessage(websocket.TextMessage, []byte(mess.Data))
	return nil
}
