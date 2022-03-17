package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/emoji"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Emoji.Flag(&emoji.FlagRequest{
		Code: "GB",
	})
	fmt.Println(rsp, err)
}
