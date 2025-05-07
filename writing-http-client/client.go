package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/handler"
	"main/model"
	"net/http"
	"os"
	"strings"
)

func main() {
	var name string
	var age int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)

	person := model.Person {
		Name: name,
		Age: age,
	}

	jsonData, err := json.Marshal(person)
	handler.HandleError(err)

	reqBody := bytes.NewBuffer(jsonData)

	resp, err := http.Post("http://127.0.0.1:8080/sendData", "application/json", reqBody)
	handler.HandleError(err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	handler.HandleError(err)

	fmt.Println("Response from server:", string(data))
}