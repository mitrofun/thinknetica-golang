package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"strings"
	"time"
)

const UrpPrefix = "ws://localhost:8080"

func send() {
	endpoint := UrpPrefix + "/send"
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("send message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.Replace(msg, "\n", "", -1)
		ws, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
		if err != nil {
			ws.Close()
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			ws.Close()
			fmt.Printf("Error sending the message: %v", err)
		}
		ws.Close()
	}
}

func messages() {
	endpoint := UrpPrefix + "/messages"
	ws, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
	}
	defer ws.Close()

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("Message could not be read: %v", err)
		}
		fmt.Printf("Message: %s.\n", p)
		time.Sleep(time.Second)
	}
}

func main() {
	go messages()
	send()
}
