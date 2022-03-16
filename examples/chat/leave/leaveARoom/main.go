package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

// Leave a chat room
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Leave(&chat.LeaveRequest{})
	fmt.Println(rsp, err)
}
