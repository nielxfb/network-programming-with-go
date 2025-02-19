package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return
	}
	//Fixed and Scanner	
	// data := []byte("String data")
	// _, err = conn.Write(data) 
	
	//Dynamically
	payload := String("Jonathan Maverick")
	_, err = payload.WriteTo(conn)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
}