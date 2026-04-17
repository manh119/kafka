package main

import (
	"log"
	"net"
)

//	Client ping and Server pong.
//
// - how to read and write from TCP stream
func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		log.Printf("read from connection %s", string(buffer[:n]))

		_, err = conn.Write([]byte("hello world from server"))
		if err != nil {
			return
		}

		conn.Close()
	}
}
