package main

import (
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Call returns a personalised "Hello $name" response
func main() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := helloworldService.Call(&helloworld.CallRequest{
		Name: "John",
	})
	fmt.Println(rsp, err)

}
