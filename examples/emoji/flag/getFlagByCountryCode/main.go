package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/emoji"
)

// Get the flag for a country. Requires country code e.g GB for great britain
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Emoji.Flag(&emoji.FlagRequest{})
	fmt.Println(rsp, err)
}
