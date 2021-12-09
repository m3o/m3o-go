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
		Id:       "user-1",
		Password: "Password1",
		Username: "joe",
	})
	fmt.Println(rsp, err)
}
