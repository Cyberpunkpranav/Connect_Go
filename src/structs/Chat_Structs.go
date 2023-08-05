package structs

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
	*websocket.Conn
}

// Response sent beck from the websockets
type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	User           string   `json:"user"`
	ConnectedUsers []string `json:"connected_users"`
}

// Websocket request from the Client
type WsPayload struct {
	Action   string              `json:"action"`
	Id       string              `json:"id"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}
