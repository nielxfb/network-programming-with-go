package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("[Server] Error: %v", err)
	}
	defer listener.Close()

	fmt.Println("[Server] Listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("[Server] Error: %v", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Printf("[Server] Accepted connection from %v\n", c.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		err := c.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Printf("[Server] Error: %v", err)
			break
		}

		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("[Server] Client %v disconnected\n", c.RemoteAddr())
				break
			}
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Printf("[Server] Read timeout from %v\n", c.RemoteAddr())
				break
			}
			log.Printf("[Server] Error: %v", err)
			break
		}

		fmt.Printf("[Server] From client %v: %v\n", c.RemoteAddr(), string(buf[:n]))
	}
}
