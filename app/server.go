package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		buff := make([]byte, 1024)
		_, err = conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading input: ", err.Error())
			continue
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}
