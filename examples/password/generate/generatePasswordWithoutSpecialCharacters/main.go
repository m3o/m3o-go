package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/password"
)

// Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Password.Generate(&password.GenerateRequest{
		Length:    16,
		Lowercase: true,
		Numbers:   true,
		Special:   false,
		Uppercase: true,
	})
	fmt.Println(rsp, err)
}
