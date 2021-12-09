package main

import (
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Subscribe to messages for a given topic.
func main() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	stream, err := mqService.Subscribe(&mq.SubscribeRequest{
		Topic: "events",
	})
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
