package controllers

import (
	structs "ConnectApp/src/structs"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/websocket"
)

var WsChan = make(chan structs.WsPayload)
var clients = make(map[structs.WebSocketConnection]string)

// Upgrade the normal connection with the websocket upgrader
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connection to the client and engaging client with websocket_Connection
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			fmt.Println("WebSocket connection going away or closed abnormally:", err)
			ws.Close()
		} else {
			log.Println("Error reading message:", err)
		}
	}

	conn := structs.WebSocketConnection{Conn: ws}
	// clients[conn] = ""
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			fmt.Println("WebSocket connection going away or closed abnormally:", err)
			ws.Close()
		} else {
			log.Println("Error reading message:", err)
		}
	}
	//Make the Routine for listenting to the websocket from the client
	go ListenForWs(&conn)
}

// Listen to the client websocket data and gather all in the channel
func ListenForWs(conn *structs.WebSocketConnection) {
	defer func() {

		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	var Payload structs.WsPayload

	for {
		err := conn.ReadJSON(&Payload)
		if err != nil {
			conn.Conn.Close()
		} else {
			Payload.Conn = *conn
			WsChan <- Payload
		}
	}
}

// Broadcast gathered data in the channel
func ListenToWsChannel() {
	var response structs.WsJsonResponse
	for {
		e := <-WsChan
		// log.Println(e.Username, e.Message, e.Action, e.Id)
		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := GetUserlist()
			response.Action = "UserLists"
			response.ConnectedUsers = users
			BroadcastToAll(response)

		case "left":
			response.Action = "UserLists"
			delete(clients, e.Conn)
			users := GetUserlist()
			response.ConnectedUsers = users
			BroadcastToAll(response)

		case "broadcast":
			response.Action = "Broadcast"
			response.User = e.Username
			response.Message = e.Message
			BroadcastToAll(response)
		}
	}
}
func GetUserlist() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}

	}
	sort.Strings(userList)
	log.Println(userList)
	return userList
}

// Broadcast gathered data in the websockets
func BroadcastToAll(response structs.WsJsonResponse) {

	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			_ = client.Close()
			delete(clients, client)
		}	
	}
}
