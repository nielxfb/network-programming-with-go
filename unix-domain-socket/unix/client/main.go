package main

import (
	"fmt"
	"net"
)

func main() {
	path := `` // Replace with socket path when server up.

	conn, err := net.Dial("unix", path)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	msg := []byte("Hello unix!")

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