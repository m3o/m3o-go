package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/helloworld"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Helloworld.Call(&helloworld.CallRequest{
		Name: "John",
	})
	fmt.Println(rsp, err)
}
