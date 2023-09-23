package websocket

import (
	"fmt"
)

type Pool struct {
	Register	chan *Client
	Unregister	chan *Client
	Clients		map[*Client]bool
	Broadcast	chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:	make(chan *Client),
		Unregister:	make(chan *Client),
		Clients:	make(map[*Client]bool),
		Broadcast:	make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register: 
			pool.Clients[client] = true
			fmt.Println("The size of the pool is ", len(pool.Clients))
			id := client.ID
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Client: id, Type: 1, Body: "New user joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("The size of the pool is ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Client: client.ID, Type: 1, Body: "User disconnected."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending a message to all clients in the pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}