package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!"))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Received: %s\n", buf[:n])
}