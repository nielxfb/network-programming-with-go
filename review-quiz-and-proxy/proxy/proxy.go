package main

import (
	"fmt"
	"io"
	"net"
)

func handleProxyConn(clientConn net.Conn){
	defer clientConn.Close()

	serverConn, err := net.Dial("tcp", "127.0.0.1:2727")
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()
	//Mengirim data server ke client
	go io.Copy(serverConn, clientConn)
	//Mengirim data client ke server
	io.Copy(clientConn, serverConn)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6666")

	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Proxy is running in port 6666")

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleProxyConn(clientConn)
	}
}
