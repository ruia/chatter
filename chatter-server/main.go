package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "2222"
	SERVER_TYPE = "tcp"

	WELCOME_MESSAGE = "Welcome to server CHATTER! Enjoy your stay %s"
)

type ClientData struct {
	name       string
	connection net.Conn
}

var clients = make(map[string]*ClientData)

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
	fmt.Print("Waiting for clients...\n\n")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting client!\n Error Msg: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Client connected.")

		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	defer func() {
		fmt.Printf("%v: Disconnected\n", connection.RemoteAddr())
		delete(clients, connection.RemoteAddr().String())
		connection.Close()
	}()

	client := new(ClientData)
	client.name = "anonymous"
	client.connection = connection

	rawDataName := getData(connection)
	fmt.Printf("received: %v", rawDataName)
	clientName := strings.Split(rawDataName, " ")
	if clientName[0] == "/setName" && clientName[1] != "" {
		client.name = strings.Trim(clientName[1], "\n")
	}

	// msg := fmt.Sprintf(WELCOME_MESSAGE, client.name)
	println(client.name)
	sendData(connection, "Welcome to server")

	clients[client.connection.RemoteAddr().String()] = client

	// reader := bufio.NewReader(connection)

	for {
		// userInput, err := reader.ReadString('\n')
		userInput, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		userInput = strings.Trim(userInput, "\n")
		// println(client.name)
		// println(userInput)
		// userInput := getData(connection)
		// fmt.Println(clients[connection.RemoteAddr().String()].name)
		println(client.name + ": " + userInput)
		broadcast(*client, userInput)
	}
}

func broadcast(currentClient ClientData, msg string) {
	for k, cd := range clients {
		if currentClient.connection.RemoteAddr().String() != k {
			sendData(cd.connection, msg)
		}
	}
}

func getData(connection net.Conn) string {
	data, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	return data
}

func sendData(connection net.Conn, data string) {
	writer := bufio.NewWriter(connection)
	_, err := writer.WriteString(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
