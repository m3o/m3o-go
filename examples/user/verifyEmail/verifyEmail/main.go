package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

// Verify the email address of an account from a token sent in an email to the user.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.VerifyEmail(&user.VerifyEmailRequest{
		Token: "012345",
	})
	fmt.Println(rsp, err)
}
