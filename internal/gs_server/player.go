package gs_server

import "github.com/gorilla/websocket"

type Player struct {
	Name    string
	Picture string
	Color   string

	connection *websocket.Conn
}
