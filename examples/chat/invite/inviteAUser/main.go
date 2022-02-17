package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Invite a user to a chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Invite(&chat.InviteRequest{})
	fmt.Println(rsp, err)
}
