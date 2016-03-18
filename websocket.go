package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	messagingService = new(messaging)
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "error upgrading to websockets", 400)
		return
	}
	defer ws.Close()

	id := messagingService.register(ws)
	defer messagingService.unregister(id)

	ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("your id is %d\n", id)))

	for {
		mt, body, err := ws.ReadMessage()
		if err != nil {
			break
		}

		switch mt {
		case websocket.TextMessage:
			m, err := validateMessage(body)
			if err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte("wrong message"))
			} else if err := messagingService.sendMessage(m); err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			}
		default:
		}
	}
}
