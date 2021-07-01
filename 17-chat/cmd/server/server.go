package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	msgQueue = make(chan string)
	mux      sync.Mutex
	clients  = make([]chan string, 0)
)

func send(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	messageType, message, err := conn.ReadMessage()
	if err != nil {
		_ = conn.WriteMessage(messageType, []byte(err.Error()))
		return
	}
	fmt.Printf("Message received: %s \n", message)

	msgQueue <- string(message)
}

func connectClient() chan string {
	mux.Lock()
	client := make(chan string)
	clients = append(clients, client)
	mux.Unlock()
	return client
}

func disconnectClient(client chan string) {
	mux.Lock()
	for i := range clients {
		if clients[i] == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
	mux.Unlock()
}

func messages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := connectClient()
	for msg := range client {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return
		}
	}

	defer conn.Close()
	defer disconnectClient(client)
}

func publishMessages() {
	for msg := range msgQueue {
		for _, c := range clients {
			c <- msg
		}
	}
}

func main() {
	http.HandleFunc("/send", send)
	http.HandleFunc("/messages", messages)

	go publishMessages()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
