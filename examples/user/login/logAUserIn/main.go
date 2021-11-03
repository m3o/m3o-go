package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Login using username or email. The response will return a new session for successful login,
// 401 in the case of login failure and 500 for any other error
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Login(&user.LoginRequest{
		Email:    "joe@example.com",
		Password: "mySecretPass123",
	})
	fmt.Println(rsp, err)

}
