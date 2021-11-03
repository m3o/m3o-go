package main

import (
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Stream returns a stream of "Hello $name" responses
func main() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := helloworldService.Stream(&helloworld.StreamRequest{
		Name: "not supported",
	})
	fmt.Println(rsp, err)
}
