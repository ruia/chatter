package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "2222"
	SERVER_TYPE = "tcp"
)

// possible data structure?
type ClientData struct {
	address    string
	name       string
	connection net.Conn
}

func main() {
	fmt.Println("Welcome to Chatter Server!")
	fmt.Println("Starting the server...")

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {
		fmt.Println("Error starting server")
		fmt.Println("Error msg:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Printf("Serving on: Host: %v - Port: %v\n", SERVER_HOST, SERVER_PORT)
	fmt.Println("Waiting for clients...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting client!\n Error Msg: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	// create a map with connected clients?
	//fmt.Println(connection.RemoteAddr())

	buffer := make([]byte, 1024)
	mLen, err := connection.Read([]byte(buffer))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	msg := string(buffer[:mLen])
	fmt.Println("Received: ", msg)
}
