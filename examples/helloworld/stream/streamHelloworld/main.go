package main

import (
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Stream returns a stream of "Hello $name" responses
func main() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	stream, err := helloworldService.Stream(&helloworld.StreamRequest{
		Messages: 10,
		Name:     "John",
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
