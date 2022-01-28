package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Update the account password
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.UpdatePassword(&user.UpdatePasswordRequest{
		ConfirmPassword: "Password2",
		NewPassword:     "Password2",
		OldPassword:     "Password1",
		UserId:          "user-1",
	})
	fmt.Println(rsp, err)
}
