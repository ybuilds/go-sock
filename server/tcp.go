package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		buffer := make([]byte, 2048)
		_, err := connection.Read(buffer)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Client: %s\n", string(buffer))

		fmt.Printf("Message: ")
		message, error := reader.ReadString('\n')

		if error != nil {
			log.Println(error)
			continue
		}

		connection.Write([]byte(message))
	}
}

func tcpServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	fmt.Println("Waiting for requests...")

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Connection accepted")

		go handleConnection(connection)
	}
}
