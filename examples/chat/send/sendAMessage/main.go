package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Connect to a chat to receive a stream of messages
// Send a message to a chat
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Send(&chat.SendRequest{
		Client:  "web",
		Subject: "Random",
		Text:    "Hey whats up?",
	})
	fmt.Println(rsp, err)
}
