package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil{
		return
	}

	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go handleClientDynamically(conn)
	}
}

func handleClientFixed(conn net.Conn){
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return
	}

	fmt.Printf("Received %s, Bytes %d\n", buffer[:n], n)
}

func handleClientScanner(conn net.Conn){
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		text := scanner.Text()
		fmt.Printf("%+v\n", text)
	}
}

func handleClientDynamically(conn net.Conn){
	defer conn.Close()

	payload, err := Decode(conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(payload)

	defer conn.Close()
}