package main

import (
	"net/http"
)

func setupRoutes() {
	// pool := websocket.NewPool()
}

func main() {
	setupRoutes()

	http.ListenAndServe(":9090", nil)
}
