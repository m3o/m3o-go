package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stream"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stream.SendMessage(&stream.SendMessageRequest{
		Channel: "general",
		Text:    "Hey checkout this tweet https://twitter.com/m3oservices/status/1455291054295498752",
	})
	fmt.Println(rsp, err)
}
