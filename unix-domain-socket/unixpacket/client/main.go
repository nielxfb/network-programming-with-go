package main

import (
	"fmt"
	"net"
)

func main() {
	path := `/tmp/streaming_server2610096310/92028.sock` // Replace with socket path when server up.

	conn, err := net.Dial("unixpacket", path)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	msg := []byte("Hello unixpacket!")

	_, err = conn.Write(msg)
	if err != nil {
		panic(err)
	}

	n, err := conn.Read(msg)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Echo from server: %s\n", msg[:n])
}