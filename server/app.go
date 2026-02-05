package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := NewServer()
	http.Handle("/socket", websocket.Handler(server.handleWebServer))
	http.Handle("/order", websocket.Handler(server.handleOrders))
	http.ListenAndServe(":3000", nil)
}
