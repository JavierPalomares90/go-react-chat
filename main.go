package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/JavierPalomares90/go-react-chat/handlers"
)

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcoming someone\n")
	io.WriteString(w, "Welcome to go-chat")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Registerring someone\n")
	io.WriteString(w, "We are registering you")
}

func handleSendMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Sending Msg\n")
	io.WriteString(w, "We received your message")

}

func main() {
	http.HandleFunc("/welcome", handlers.HandleWelcome)
	http.HandleFunc("/register", handlers.HandleRegister)
	http.HandleFunc("/sendMsg", handlers.HandleSendMsg)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}
}
