package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Reset password with the code sent by the "SendPasswordResetEmail" endoint.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.ResetPassword(&user.ResetPasswordRequest{
		Code:            "012345",
		ConfirmPassword: "NewPassword1",
		Email:           "joe@example.com",
		NewPassword:     "NewPassword1",
	})
	fmt.Println(rsp, err)
}
