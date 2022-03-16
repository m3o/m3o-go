package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/email"
)

// Parse an RFC5322 address e.g "Joe Blogs <joe@example.com>"
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Email.Parse(&email.ParseRequest{
		Address: "Joe Blogs <joe@example.com>",
	})
	fmt.Println(rsp, err)
}
