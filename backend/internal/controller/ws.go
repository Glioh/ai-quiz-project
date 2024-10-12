package controller

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

type WebsocketController struct{}

func Ws() WebsocketController {
	return WebsocketController{}
}

func (c WebsocketController) Ws(conn *websocket.Conn) {
	fmt.Println("New WebSocket connection established")

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received message: %s\n", string(msg))

		if err := conn.WriteMessage(mt, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

	fmt.Println("WebSocket connection closed")
}
