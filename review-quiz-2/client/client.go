package client

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"main/handler"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkFriends() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/vk/friends", nil)
	handler.HandleError(err)

	resp, err := http.DefaultClient.Do(req)
	handler.HandleError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handler.HandleError(err)

	fmt.Println("Response from server:", string(body))
}

func sendMessage() {
	var name string
	var msg string

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
		re := regexp.MustCompile(`^[a-zA-Z]+$`)
		if re.MatchString(name) {
			break
		}
	}

	for {
		fmt.Print("Enter your message: ")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if len(msg) >= 9 {
			break
		}
	}

	reqBody := new(bytes.Buffer)
	w := multipart.NewWriter(reqBody)

	file, err := os.Create(name + ".txt")
	handler.HandleError(err)

	defer file.Close()
	
	_, err = file.WriteString(msg)
	handler.HandleError(err)

	part, err := w.CreateFormFile("file", name+".txt")
	handler.HandleError(err)

	_, err = io.Copy(part, file)
	handler.HandleError(err)

	w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://127.0.0.1:8080/vk/send", reqBody)
	handler.HandleError(err)

	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	handler.HandleError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handler.HandleError(err)

	fmt.Println("Response from server:", string(body))
}

func Serve() {
	for {
		fmt.Println("Welcome to VK.com")
		fmt.Println("1. Check friends")
		fmt.Println("2. Send a message")
		fmt.Println("0. Exit")
		fmt.Print(">> ")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			checkFriends()
		} else if choice == 2 {
			sendMessage()
		} else if choice == 0 {
			break
		}
	}
}