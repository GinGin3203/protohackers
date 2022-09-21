package main

import (
	"fmt"
	"net"
)

const (
	HOST = "localhost"
	PORT = "1025"
)

func main() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		fmt.Println(conn)
		if err != nil {
			panic(err)
		}
		go func(c *net.TCPConn) {
			buffer := make([]byte, 512)
			c.Read(buffer)
			fmt.Println(buffer)
			c.Write(buffer)
		}(conn)
	}
}
