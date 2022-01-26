package main

import (
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Validate an email address format
func main() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Validate(&email.ValidateRequest{
		Address: "joe@example.com",
	})
	fmt.Println(rsp, err)
}
