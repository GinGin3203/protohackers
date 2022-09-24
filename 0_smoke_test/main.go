package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	host = "0.0.0.0"
	port = "1025"
)

func main() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", host+":"+port)
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
			defer c.Close()
			for {
			bytes, err := io.ReadAll(c)
			if err != nil {
				log.Println(err)		
				return
			}			
			if len(bytes) == 0{
				continue
			}
			fmt.Println(bytes)
			if err != nil {
				log.Println(err.Error())
				return
			}
			c.Write(bytes)
			}
		}(conn)
	}
}
