package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// List the messages in a chat
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.History(&chat.HistoryRequest{})
	fmt.Println(rsp, err)
}
