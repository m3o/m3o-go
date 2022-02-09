package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Send a verification email to a user.
// Email "from" will be 'noreply@email.m3ocontent.com'.
// The verification link will be injected in the email
// as a template variable, $micro_verification_link e.g
// 'Welcome to M3O! Use the link below to verify your email: $micro_verification_link'
// The variable will be replaced with a url similar to:
// 'https://user.m3o.com/user/verify?token=a-verification-token&redirectUrl=your-redir-url'
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendVerificationEmail(&user.SendVerificationEmailRequest{
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
