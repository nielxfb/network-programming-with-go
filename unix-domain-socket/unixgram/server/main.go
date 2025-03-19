package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

func main() {
	dir, err := ioutil.TempDir("", "datagram_server")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Directory: %s\n", dir)

	socketPath := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	fmt.Printf("Socket: %s\n", socketPath)

	defer func() {
		err = os.RemoveAll(dir)
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.ListenPacket("unixgram", socketPath)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// CHMOD 622
	err = os.Chmod(socketPath, os.ModeSocket|0622)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		_, err = conn.WriteTo(buf[:n], addr)
		if err != nil {
			panic(err)
		}
	}
}
