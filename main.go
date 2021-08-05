package main

import (
	"log"
	"net/http"
	"websocket/server"
	//"video-chat-app/server"
	//"websocket/server"
)

func main() {
	server.ALLRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting Server on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
