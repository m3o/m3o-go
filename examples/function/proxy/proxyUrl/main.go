package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// Return the backend url for proxying
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Proxy(&function.ProxyRequest{
		Id: "helloworld",
	})
	fmt.Println(rsp, err)
}
