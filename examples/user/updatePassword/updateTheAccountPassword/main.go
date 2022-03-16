package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

// Update the account password
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.UpdatePassword(&user.UpdatePasswordRequest{
		ConfirmPassword: "Password2",
		NewPassword:     "Password2",
		OldPassword:     "Password1",
		UserId:          "user-1",
	})
	fmt.Println(rsp, err)
}
