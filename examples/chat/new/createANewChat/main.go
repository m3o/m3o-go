package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Create a new chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.New(&chat.NewRequest{
		Description: "The general chat room",
		Name:        "general",
	})
	fmt.Println(rsp, err)
}
