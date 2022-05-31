package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

const (
	SERVER_TYPE = "tcp"
)

var delimiter byte
var delimiterFull string

func main() {
	oper := runtime.GOOS
	switch oper {
	case "windows":
		delimiter = '\r'
		delimiterFull = "\r\n"
	default:
		delimiter = '\n'
		delimiterFull = "\n"
	}

	// TODO later read args from command line?
	serverHost := "127.0.0.1"
	serverPort := "2222"

	fmt.Println("Welcome to Chatter Client!")

	fmt.Println("Please enter your name: ")
	name := ""
	for name == "" {
		name = readInput()
	}

	fmt.Println("Connecting to server...")

	connection, err := net.Dial(SERVER_TYPE, serverHost+":"+serverPort)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to: %v:%v\n\n", serverHost, serverPort)
	sendData(connection, "/setName "+name)

	go handleDataFromServer(connection)

	input := ""
	for input != "/exit" {
		input = readInput()

		if input != "/exit" {
			sendData(connection, input)
		}
	}
}

func handleDataFromServer(connection net.Conn) {
	reader := bufio.NewReader(connection)
	for {
		data, err := reader.ReadString(delimiter)
		data = strings.Trim(data, delimiterFull)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		fmt.Println(">>", data)
	}
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

func readInput() string {
	input, err := bufio.NewReader(os.Stdin).ReadString(delimiter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return input
}
