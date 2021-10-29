package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/user"
)

// Logout a user account
func LogAuserOut() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Logout(&user.LogoutRequest{
		SessionId: "sds34s34s34-s34s34-s43s43s34-s4s34s",
	})
	fmt.Println(rsp, err)
}
