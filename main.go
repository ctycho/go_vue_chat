package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/ctycho/go_vue_chat/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endoint reached")

	conn, err := websocket.Upgrade(w, r)
	if (err != nil) {
		fmt.Printf("UPGRADE %+v\n", err)
	}

	id := strconv.Itoa(rand.Intn(100))
	fmt.Println(id)
	client := &websocket.Client{
		ID: id,
		Conn: conn, 
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat has started")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}