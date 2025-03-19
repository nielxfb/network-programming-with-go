package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	tempDir := ``
	tempSocket := ``

	serverAddr, err := net.ResolveUnixAddr("unixgram", tempSocket)
	if err != nil {
		panic(err)
	}

	clientSocket := filepath.Join(tempDir, fmt.Sprintf("c%d.sock", os.Getpid()))

	conn, err := net.ListenPacket("unixgram", clientSocket)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = os.Chmod(clientSocket, os.ModeSocket|0622)
	if err != nil {
		panic(err)
	}

	msg := []byte("Hello unixgram!")

	_, err = conn.WriteTo(msg, serverAddr)
	if err != nil {
		panic(err)
	}

	n, _, err := conn.ReadFrom(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Echo from server: %s\n", msg[:n])
}
