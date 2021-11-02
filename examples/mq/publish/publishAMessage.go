package example

import (
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Publish a message. Specify a topic to group messages for a specific topic.
func PublishAmessage() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := mqService.Publish(&mq.PublishRequest{
		Message: map[string]interface{}{
			"id":   "1",
			"type": "signup",
			"user": "john",
		},
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
