package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Leave a chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Leave(&chat.LeaveRequest{})
	fmt.Println(rsp, err)
}
