package websocket

import (
	"log"
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID		string
	Conn	*websocket.Conn
	Pool	*Pool
	mu		sync.Mutex
}

type Message struct {
	Client	string`json:"client"`
	Type	int `json:"type"`
	Body	string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		message := Message{Client: c.ID, Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
	}
}