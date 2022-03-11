package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Logout of all user's sessions
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.LogoutAll(&user.LogoutAllRequest{})
	fmt.Println(rsp, err)
}
