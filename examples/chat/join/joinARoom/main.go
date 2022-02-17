package main

import (
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Join a chat room
func main() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	stream, err := chatService.Join(&chat.JoinRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
