package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.SendVerificationEmail(&user.SendVerificationEmailRequest{
		Email:              "joe@example.com",
		FailureRedirectUrl: "https://m3o.com/verification-failed",
		FromName:           "Awesome Dot Com",
		RedirectUrl:        "https://m3o.com",
		Subject:            "Email verification",
		TextContent: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,
	})
	fmt.Println(rsp, err)
}
