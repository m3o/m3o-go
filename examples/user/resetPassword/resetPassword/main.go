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
		Code:            "some-code-from-email",
		ConfirmPassword: "newpass123",
		NewPassword:     "newpass123",
	})
	fmt.Println(rsp, err)

}
