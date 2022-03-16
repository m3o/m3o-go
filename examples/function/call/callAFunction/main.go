package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// Call a function by name
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Call(&function.CallRequest{
		Name: "helloworld",
		Request: map[string]interface{}{
			"name": "Alice",
		},
	})
	fmt.Println(rsp, err)
}
