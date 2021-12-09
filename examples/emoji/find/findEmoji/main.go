package main

import (
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Find an emoji by its alias e.g :beer:
func main() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Find(&emoji.FindRequest{
		Alias: ":beer:",
	})
	fmt.Println(rsp, err)
}
