package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2727")

	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server is running in port 2727")

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleServerConn(clientConn)
	}
}

func handleServerConn(clientConn net.Conn) {
	defer clientConn.Close()

	//Read
	var size uint32
	err := binary.Read(clientConn, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size) //Buat buffer dengan type data []slice dan ukuran nya size
	clientConn.Read(buffer)      //Baca clientConn, dan masukkan ke buffer
	receive := string(buffer)    //Ubah menjadi string

	fmt.Println("Server RECEIVED: " + receive)

	// time.Sleep(5 * time.Second)
	//Write balik ke client nya
	var response string
	if receive == "I hate netvork!" {
		response = "I hate netvork toooooooooo!"
	} else {
		response = "Your message: " + receive
	}

	binary.Write(clientConn, binary.BigEndian, uint32(len(response)))
	clientConn.Write([]byte(response))
}
