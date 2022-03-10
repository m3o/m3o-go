package main

import (
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Publish a message. Specify a topic to group messages for a specific topic.
func main() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := mqService.Publish(&mq.PublishRequest{
		Message: map[string]interface{}{
			"user": "john",
			"id":   "1",
			"type": "signup",
		},
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
