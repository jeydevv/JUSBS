package main

// Import modules
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Constant variables declarations
const (
	PROTOCOL = "tcp"
	PORT     = ":1111" // Enter same port as in the client, instead of 1111
)

func main() {
	fmt.Println("Starting server~")

	// Open listener and connection to client
	for {
		listener, _ := net.Listen(PROTOCOL, PORT)
		connection, _ := listener.Accept()

		message, _ := bufio.NewReader(connection).ReadString('|')
		fmt.Print("Recieved: ", string(message))

		// Close listener and connection
		listener.Close()
		connection.Close()

		// Write info to file
		WriteToFile(message)
	}
}

// WriteToFile writes recieved buffer to a file
func WriteToFile(message string) {
	file, _ := os.Create(message[0:strings.Index(message, "\n")] + ":-log.txt")
	file.WriteString(message)
}
