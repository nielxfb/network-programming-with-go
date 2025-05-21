package main

import (
	"fmt"
	"main/client"
	"main/server"
)

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Start Server")
		fmt.Println("2. Start Client")
		fmt.Println("0. Exit")
		fmt.Print(">> ")
		var choice int
		fmt.Scanln(&choice)
		if choice == 0 {
			break
		} else if choice == 1 {
			server.Serve()
		} else if choice == 2 {
			client.Serve()
		}
	}
}