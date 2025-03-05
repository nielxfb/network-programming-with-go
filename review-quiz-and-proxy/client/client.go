package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Main Menu NetVorK")
		fmt.Println("1. Input")
		fmt.Println("2. Exit")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			input()
		} else if choice == 2 {
			fmt.Println("Exited from Program")
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	var message string

	for {
		fmt.Print("Enter your message [at least 6 character and ends with '!']: ")
		scanner.Scan()
		message = scanner.Text()

		if len(message) <= 6 {
			fmt.Println("Message too short")
		} else if message[len(message)-1] != '!' {
			fmt.Println("Message must ends with !")
		} else {
			break
		}
	}
	sendToServer(message)
}

func sendToServer(message string) {
	dial, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}

	//Write
	binary.Write(dial, binary.BigEndian, uint32(len(message)))
	_, err = dial.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	//Implement Deadline -> 5 POINT
	dial.SetReadDeadline(time.Now().Add(3 * time.Second))

	//Read
	var size uint32
	err = binary.Read(dial, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size) //Buat buffer dengan type data []slice dan ukuran nya size
	dial.Read(buffer)            //Baca clientConn, dan masukkan ke buffer
	receive := string(buffer)    //Ubah menjadi string
	fmt.Println(receive)

	defer dial.Close()
}
