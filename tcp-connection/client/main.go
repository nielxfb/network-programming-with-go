package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	fmt.Println("[Client] Connected to server")

	for {
		fmt.Print("[Client] Input message: ")
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')

		msg = msg[:len(msg)-1]

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("[Client] Error: %v", err)
		}

		fmt.Println("[Client] Message sent")
	}
}