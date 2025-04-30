package main

import (
	"fmt"
	"net/http"

	"github.com/example.com/handler"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// Dia ini pakai method get atau bukan
func helloMiddleware(method string,
	handlerFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Misalkan method tidak sesuai
		if r.Method != method {
			http.Error(w, "Method Not Allowed",
				http.StatusMethodNotAllowed)
			return
		}
		// Misalkan method nya sesuai
		handlerFunc(w, r)
	}

}

func main() {
	addr := "localhost:1234"
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloMiddleware(http.MethodGet, helloHandler))

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Println("Server is running on", addr)
	// HTTP
	// err := server.ListenAndServe()

	// HTTPS
	err := server.ListenAndServeTLS("127.0.0.1.pem", "127.0.0.1-key.pem")
	handler.HandleError(err)
}
