package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

// Kick a user from a chat room
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Kick(&chat.KickRequest{})
	fmt.Println(rsp, err)
}
