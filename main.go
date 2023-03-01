package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = ":8080"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", handle)

	fmt.Printf("listen and serve: %s\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			panic(err)
		}
		fmt.Printf("client send messageType:%d, message:%s\n", messageType, string(p))
	}
}
