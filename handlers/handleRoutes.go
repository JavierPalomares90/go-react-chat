package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("In my module\n")
	fmt.Printf("Welcoming someone\n")
	io.WriteString(w, "Welcome to go-chat")
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Registerring someone\n")
	io.WriteString(w, "We are registering you")
}

func HandleSendMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Sending Msg\n")
	io.WriteString(w, "We received your message")

}
