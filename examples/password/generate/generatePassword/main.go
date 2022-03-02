package main

import (
	"fmt"
	"os"

	"go.m3o.com/password"
)

// Generate a strong random password
func main() {
	passwordService := password.NewPasswordService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := passwordService.Generate(&password.GenerateRequest{
		Length: 16,
	})
	fmt.Println(rsp, err)
}
