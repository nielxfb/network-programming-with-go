package main

import (
	"encoding/json"
	"fmt"
	"io"
	"main/handler"
	"main/model"
	"net/http"
)

func handleSendData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	handler.HandleError(err)
	defer r.Body.Close()

	var person model.Person
	err = json.Unmarshal(body, &person)

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Fprintln(w, "Data received successfully")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/sendData", handleSendData)

	server := http.Server {
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	handler.HandleError(err)
}