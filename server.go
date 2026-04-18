package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	stream := bufio.NewReadWriter(
		bufio.NewReader(conn),
		bufio.NewWriter(conn),
	)

	header, err := stream.ReadByte()
	if err != nil {
		return
	}

	data := make([]byte, header)
	_, err = stream.Read(data)
	if err != nil {
		return
	}

	fmt.Printf("Data: %s\n", data)

	stream.WriteString("6hellos")
	stream.Flush()
}

//
//func main() {
//	listener, _ := net.Listen("tcp", ":8081")
//
//	for {
//		conn, _ := listener.Accept()
//		go handle(conn)
//	}
//}
