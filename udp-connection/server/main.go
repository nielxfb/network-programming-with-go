package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", "127.0.0.1:8080")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	fmt.Println("Server listening on port 8080")

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Received: %s from %s\n", buf[:n], addr)

		_, err = conn.WriteTo([]byte("Hello, client!"), addr)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Sent: echo reply to %s\n", addr)
	}
}
