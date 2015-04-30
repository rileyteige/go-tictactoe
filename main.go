// server project main.go
package main

import (
	"fmt"
	"github.com/rileyteige/go-tictactoe/server"
	"log"
)

const port = 8081

func main() {
	fmt.Printf("Starting Tic-Tac-Toe server on port %v...\n", port)

	err := server.Listen(port)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
