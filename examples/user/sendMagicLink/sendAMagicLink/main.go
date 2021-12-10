package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Login using email only - Passwordless
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	stream, err := userService.SendMagicLink(&user.SendMagicLinkRequest{
		Address:  "www.example.com",
		Email:    "joe@example.com",
		Endpoint: "verifytoken",
		FromName: "Awesome Dot Com",
		Subject:  "MagicLink to access your account",
		TextContent: `Hi there,

Click here to access your account $micro_verification_link`,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
