package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

func main() {
	dir, err := ioutil.TempDir("", "streaming_server")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Directory: %s\n", dir)

	socketPath := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	fmt.Printf("Socket: %s\n", socketPath)
	fmt.Println("Server up!")

	defer func() {
		err = os.RemoveAll(dir)
		if err != nil {
			panic(err)
		}
	}()

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		panic(err)
	}

	_, err = c.Write(buf[:n])
	if err != nil {
		panic(err)
	}
}
