package handlers

import (
	"fmt"
	"io"
	"net/http"
	"github.com/JavierPalomares90/go-react-chat/user"

)

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcoming someone\n")
	io.WriteString(w, "Welcome to go-chat. You'll need to register to send a message")
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Registerring someone\n")
	io.WriteString(w, "We are registering you")
	//TODO: Pass in the appropriate methods to regsiter the user
	user.RegisterUser()
	// Get the user name, generate a JWT Token, return it to the user
	/**
	 ** TODO: Complete impl
	 */
}

func HandleSendMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Sending Msg\n")
	io.WriteString(w, "We received your message")

	/**
	 ** TODO: Complete impl
	 */
}
