package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/spam"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Spam.Classify(&spam.ClassifyRequest{
		From:    "noreply@m3o.com",
		Subject: "Welcome",
		To:      "hello@example.com",
	})
	fmt.Println(rsp, err)
}
