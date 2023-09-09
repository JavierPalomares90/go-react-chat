package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcoming someone\n")
	io.WriteString(w, "Welcome to go-chat")
}

func main() {
	http.HandleFunc("/welcome",handleWelcome)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}
}