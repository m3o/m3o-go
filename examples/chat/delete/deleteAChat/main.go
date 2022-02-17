package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Delete a chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Delete(&chat.DeleteRequest{})
	fmt.Println(rsp, err)
}
