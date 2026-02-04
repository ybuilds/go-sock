package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func tcpClient() {
	scanner := bufio.NewReader(os.Stdin)

	connection, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
		return
	}

	defer connection.Close()

	for {
		fmt.Print("Message: ")
		message, err := scanner.ReadString('\n')

		if err != nil {
			log.Println(err)
			continue
		}

		_, err = connection.Write([]byte(message))

		if err != nil {
			log.Println(err)
			continue
		}

		buffer := make([]byte, 2048)

		_, err = connection.Read(buffer)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Server: %s\n", string(buffer))
	}
}
