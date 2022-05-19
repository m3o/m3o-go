package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/mq"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Mq.Publish(&mq.PublishRequest{
		Message: map[string]interface{}{
			"type": "signup",
			"user": "john",
			"id":   "1",
		},
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
