package server

import (
	"fmt"
	"net/http"
)

func middleware(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			// w.WriteHeader(405)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func getFriendsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have no friends.. :(")
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Message successfully received with attached file.")
}

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/vk/friends", middleware(http.MethodGet, getFriendsHandler))
	mux.HandleFunc("/vk/send", middleware(http.MethodPost, sendMessageHandler))

	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	defer server.Close()
	server.ListenAndServe()
}