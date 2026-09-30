package ws

import "github.com/gorilla/websocket"

// Abstract struct between websocket connection and the hub
type Client struct {
	hub  *Hub
	ws   *websocket.Conn
	send chan []byte // not used yet
}
