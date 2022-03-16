package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/chat"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Chat.Create(&chat.CreateRequest{
		Description: "The general chat room",
		Name:        "general",
	})
	fmt.Println(rsp, err)
}
