package main

import (
	"fmt"
	"main/server"
	"main/client"
)

func main() {
	for {
		var choice int
		fmt.Println("Choose an option:")
		fmt.Println("1. Server")
		fmt.Println("2. Client")
		fmt.Println("0. Exit")
		fmt.Print(">> ")
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		} else if choice == 1 {
			server.Run()
		} else if choice == 2 {
			client.Run()
		}
	}
}