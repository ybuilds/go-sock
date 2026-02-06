package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func clientSocket(url, origin string) {
	socket, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatalln("Error creating client socket:", err)
		return
	}

	buffer := make([]byte, 2048)
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Message: ")
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Error reading client message:", err)
			continue
		}

		socket.Write([]byte(message))

		n, err := socket.Read(buffer)

		if err != nil {
			log.Println("Error reading from server:", err)
			continue
		}

		fmt.Println("Server:", string(buffer[:n]))
	}
}
