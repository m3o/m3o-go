package main

import (
	"fmt"
	"os"

	"go.m3o.com/email"
)

// Send an email by passing in from, to, subject, and a text or html body
func main() {
	emailService := email.NewEmailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emailService.Send(&email.SendRequest{
		From:    "Awesome Dot Com",
		Subject: "Email verification",
		TextBody: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,
	})
	fmt.Println(rsp, err)
}
