// server project main.go
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveEcho(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request.")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	echo := r.URL.Query().Get("txt")
	if echo == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, echo)
}

func main() {
	fmt.Println("Starting Tic-Tac-Toe server:")

	http.HandleFunc("/echo", serveEcho)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
