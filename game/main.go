// server project main.go
package main

import (
	"fmt"
	"log"
	"server"
)

const port = 8080

func main() {
	fmt.Printf("Starting Tic-Tac-Toe server on port %v...\n", port)

	err := server.Listen(port)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
