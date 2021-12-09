package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Update the account username or email
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Update(&user.UpdateRequest{
		Email: "joe+2@example.com",
		Id:    "user-1",
	})
	fmt.Println(rsp, err)
}
