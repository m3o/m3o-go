package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Kick a user from a chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Kick(&chat.KickRequest{})
	fmt.Println(rsp, err)
}
