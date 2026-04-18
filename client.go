package main

//
//func main() {
//	conn, _ := net.Dial("tcp", "localhost:8081")
//
//	stream := bufio.NewReadWriter(
//		bufio.NewReader(conn),
//		bufio.NewWriter(conn),
//	)
//
//	stream.WriteString("5hello")
//	stream.Flush()
//
//	header, _ := stream.ReadByte()
//
//	data := make([]byte, header)
//	stream.Read(data)
//
//	log.Printf("Got: %s", data)
//}
