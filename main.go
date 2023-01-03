package main

import (
	"net/http"

	"github.com/bagusshndr/chat-golang-api/pkg/websocket"
)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
}

func main() {
	setupRoutes()

	http.ListenAndServe(":9090", nil)
}
