package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/emoji"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Emoji.Print(&emoji.PrintRequest{
		Text: "let's grab a :beer:",
	})
	fmt.Println(rsp, err)
}
