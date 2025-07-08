package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/utils"
	"github.com/codecrafters-io/redis-starter-go/app/config"
)
var ClusterConfig := config.NewClusterInfo()

var registry = commands.NewCommandHandler()

type ClientConn struct{
	conn net.Conn
}

func NewClientConn(conn net.Conn) *ClientConn{
	return &ClientConn{conn: conn}
}

func (c *ClientConn) Read() ([]byte, error) {
	buff := make([]byte, 1024)
	n, err := c.conn.Read(buff)
	if err != nil {
		return nil, err
	}
	return buff[:n], nil
}

func (c *ClientConn) Write(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

func handleClient(conn net.Conn){
	defer conn.Close()
	client := NewClientConn(conn)

	for{
		data, err := client.Read()
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Println("Received data:", string(data))

		
		parsedData, err := utils.ParseRESP(string(data))
		if err != nil {
			fmt.Println("Error parsing data:", err)
			return
		}
		fmt.Println("Parsed data:", parsedData)

		var response string
		if parsedArray, ok := parsedData.([]interface{}); ok && len(parsedArray) > 0 {
			if cmdName, ok := parsedArray[0].(string); ok {
				args := make([]string, len(parsedArray)-1)
				for i := 1; i < len(parsedArray); i++ {
					if arg, ok := parsedArray[i].(string); ok {
						args[i-1] = arg
					}
				}
				response, err = registry.ExecuteCommand(cmdName, args)
				if err != nil {
					fmt.Println("Error executing command:", err)
					response = "-ERR " + err.Error() + "\r\n"
				}
			} else {
				response = "-ERR Invalid command format\r\n"
			}
		} else {
			response = "-ERR Invalid command format\r\n"
		}

		err = client.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func main(){
	fmt.Println("Logs from your program will appear here!")
	var port int 
	flag.IntVar(&port, "port", 6379, "Port to listen on")
	flag.Parse()

	port:=":"+strconv.Itoa(port)
	config.Initialize(port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening on port ", port, ":", err)
		os.Exit(1)
	}

	fmt.Println("Server started on port ", port)

	for{
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn)
	}
}
