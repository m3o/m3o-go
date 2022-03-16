package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/event"
)

// Read stored events
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Event.Read(&event.ReadRequest{
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
