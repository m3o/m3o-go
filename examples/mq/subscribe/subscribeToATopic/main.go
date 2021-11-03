package main

import (
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Subscribe to messages for a given topic.
func main() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := mqService.Subscribe(&mq.SubscribeRequest{
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
