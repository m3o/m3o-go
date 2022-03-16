package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Send(&chat.SendRequest{
		Client:  "web",
		Subject: "Random",
		Text:    "Hey whats up?",
	})
	fmt.Println(rsp, err)
}
