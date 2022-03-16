package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/user"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.User.SendMagicLink(&user.SendMagicLinkRequest{
		Address:  "www.example.com",
		Email:    "joe@example.com",
		Endpoint: "verifytoken",
		FromName: "Awesome Dot Com",
		Subject:  "MagicLink to access your account",
		TextContent: `Hi there,

Click here to access your account $micro_verification_link`,
	})
	fmt.Println(rsp, err)
}
