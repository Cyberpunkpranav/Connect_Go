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
		log.Println(err)
	}
	log.Println("Client Successfully Connected")
	conn := structs.WebSocketConnection{Conn: ws}
	clients[conn] = ""
	if err != nil {
		log.Println(err)
	}
	//Make the Routine for listenting to the websocket from the client
	go ListenForWs(&conn)
}

// Listen to the client websocket data and gather all in the channel
func ListenForWs(conn *structs.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("Error:%v", r))
		}
	}()
	var Payload structs.WsPayload
	for {
		err := conn.ReadJSON(&Payload)
		if err != nil {
			log.Println(err)
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
		log.Println(e.Action)
		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := GetUserlist()
			log.Println(users)
			response.Action = "UserLists"
			response.ConnectedUsers = users
			BroadcastToAll(response)

		case "left":
			response.Action = "UserList"
			delete(clients, e.Conn)
			users := GetUserlist()
			response.ConnectedUsers = users
			BroadcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("%s:%s", e.Username, e.Message)
			BroadcastToAll(response)
		}
	}
}
func GetUserlist() []string {
	var userList []string
	log.Println(clients)
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}

	}
	sort.Strings(userList)
	return userList
}

// Broadcast gathered data in the websockets
func BroadcastToAll(response structs.WsJsonResponse) {

	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
