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

	connArr := make([]net.Conn, 0)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				os.Exit(1)
			}
			connArr = append(connArr, conn)
		}
	}()

	for {
		for _, c := range connArr {
			buff := make([]byte, 1024)
			_, err = c.Read(buff)
			if err != nil {
				fmt.Println("Error reading input: ", err.Error())
				continue
			}

			c.Write([]byte("+PONG\r\n"))
		}
	}
}
