package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/email"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Email.Validate(&email.ValidateRequest{
		Address: "joe@example.com",
	})
	fmt.Println(rsp, err)
}
