package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

// List the messages in a chat
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.History(&chat.HistoryRequest{})
	fmt.Println(rsp, err)
}
