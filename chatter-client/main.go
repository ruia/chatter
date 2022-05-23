package main

import (
	"fmt"
	"net"
)

const (
	SERVER_TYPE = "tcp"
)

func main() {
	// TODO later read from comand line args?
	serverHost := "127.0.0.1"
	serverPort := "2222"

	fmt.Println("Welcome to Chatter Client!")
	fmt.Println("Connecting to server...")

	connection, err := net.Dial(SERVER_TYPE, serverHost+":"+serverPort)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to: %v:%v\n\n", serverHost, serverPort)

	// _, err = connection.Write([]byte("Hello Server! Greetings."))

	input := ""
	for input != "/exit" {
		fmt.Print("Input: ")
		fmt.Scan(&input)

		if input != "/exit" {
			_, err = connection.Write([]byte(input))
		}
	}

}
