package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{make(map[*websocket.Conn]bool)}
}

func (server *Server) handleWebServer(socket *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", socket.RemoteAddr())

	server.conns[socket] = true
	server.readLoop(socket)
}

func (server *Server) readLoop(socket *websocket.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := socket.Read(buffer)

		if err != nil {
			if err == io.EOF {
				log.Println("Client closed connection.", err)
				break
			}
			log.Println("Error reading from socket.", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Println("Client:", message)

		socket.Write([]byte("Acknowledged message\n"))

		server.broadcast([]byte(message))
	}
}

func (server *Server) broadcast(b []byte) {
	for ws := range server.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println("Write error:", err)
			}
		}(ws)
	}
}

func (server *Server) handleOrders(socket *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", socket.RemoteAddr())
	for {
		message := fmt.Sprint("subscribed and ordered:", time.Now().UnixNano())
		socket.Write([]byte(message))
		time.Sleep(time.Second * 2)
	}
}
