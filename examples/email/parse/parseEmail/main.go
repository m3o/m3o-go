package main

import (
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Parse an RFC5322 address e.g "Joe Blogs <joe@example.com>"
func main() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Parse(&email.ParseRequest{
		Address: "Joe Blogs <joe@example.com>",
	})
	fmt.Println(rsp, err)
}
