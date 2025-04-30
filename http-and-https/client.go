package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/example.com/handler"
)

func main() {
	// GET -> Mengambil data
	// POST -> Mengirim data ke server
	// PUT -> Melakukan update
	// DELETE -> Menghapus data

	// r := strings.NewReader("Hello, Reader!")

	resp, err := http.Get("https://127.0.0.1:1234/hello")
	handler.HandleError(err)

	data, err := io.ReadAll(resp.Body)
	handler.HandleError(err)
	fmt.Println("Server response:", string(data))
}
