package server

import (
	"fmt"
	"net/http"
)

func validateMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func getFriendsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have no friends.. :(")
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Message successfully received with attached file.")
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/vk/friends", validateMethod(http.MethodGet, getFriendsHandler))
	mux.HandleFunc("/vk/send", validateMethod(http.MethodPost, sendHandler))

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
