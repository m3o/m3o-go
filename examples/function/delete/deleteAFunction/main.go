package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// Delete a function by name
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Delete(&function.DeleteRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
