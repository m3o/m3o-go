package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/mq"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	stream, err := client.Mq.Subscribe(&mq.SubscribeRequest{
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
