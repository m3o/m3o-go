package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Verify the email address of an account from a token sent in an email to the user.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.VerifyEmail(&user.VerifyEmailRequest{
		Email: "joe@example.com",
		Token: "t2323t232t",
	})
	fmt.Println(rsp, err)

}
