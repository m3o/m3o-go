package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.ResetPassword(&user.ResetPasswordRequest{
		Code:            "012345",
		ConfirmPassword: "NewPassword1",
		Email:           "joe@example.com",
		NewPassword:     "NewPassword1",
	})
	fmt.Println(rsp, err)
}
