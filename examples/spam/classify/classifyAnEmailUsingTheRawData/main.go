package main

import (
	"fmt"
	"os"

	"go.m3o.com/spam"
)

// Check whether an email is likely to be spam based on its attributes
func main() {
	spamService := spam.NewSpamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spamService.Classify(&spam.ClassifyRequest{})
	fmt.Println(rsp, err)
}
