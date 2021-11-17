package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Logout a user account
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Logout(&user.LogoutRequest{
		SessionId: "df91a612-5b24-4634-99ff-240220ab8f55",
	})
	fmt.Println(rsp, err)

}
