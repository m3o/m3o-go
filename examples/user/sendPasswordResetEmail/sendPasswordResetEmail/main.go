package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Send an email with a verification code to reset password.
// Call "ResetPassword" endpoint once user provides the code.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendPasswordResetEmail(&user.SendPasswordResetEmailRequest{
		Email:    "joe@example.com",
		FromName: "Awesome Dot Com",
		Subject:  "Password reset",
		TextContent: `Hi there,
 click here to reset your password: myapp.com/reset/code?=$code`,
	})
	fmt.Println(rsp, err)
}
