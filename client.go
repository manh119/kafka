package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	for {
		log.Printf("Connecting to %s", conn.RemoteAddr())
		conn.Write([]byte("Send hello world to the server"))

		// todo : how buffer really work in Read func ?
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		log.Printf("Read from server %s", string(buffer[:n]))
		conn.Close()
	}

}
