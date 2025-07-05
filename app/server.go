package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)


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
	client:=NewClientConn(conn)

	for{
		data,err:=client.Read()
		if err!= nil{
			fmt.Println("Error reading from connection:",err)
			return
		}
		fmt.Println("Received data:",string(data))

		err=client.Write([]byte("+PONG\r\n"))
		if err!= nil{
			fmt.Println("Error writing to connection:",err)
			return
		}
	}
}

func main(){
	fmt.Println("Logs from your program will appear here!")

	l,err:=net.Listen("tcp",":6379")
	if err!= nil{
		fmt.Println("Error listening on port 6379:",err)
		os.Exit(1)
	}

	fmt.Println("Server started on port 6379")

	for{
		conn,err:=l.Accept()
		if err!= nil{
			fmt.Println("Error accepting connection:",err)
			continue
		}
		go handleClient(conn)
	}
}