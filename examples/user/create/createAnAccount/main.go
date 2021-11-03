package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Create a new user account. The email address and username for the account must be unique.
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Create(&user.CreateRequest{
		Email:    "joe@example.com",
		Id:       "usrid-1",
		Password: "mySecretPass123",
		Username: "usrname-1",
	})
	fmt.Println(rsp, err)
}
