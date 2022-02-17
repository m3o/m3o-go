package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// List available chats
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.List(&chat.ListRequest{})
	fmt.Println(rsp, err)
}
