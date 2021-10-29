package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/emoji"
)

// Send an emoji to anyone via SMS. Messages are sent in the form '<message> Sent from <from>'
func SendAtextContainingAnEmojiToAnyoneViaSms() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Send(&emoji.SendRequest{
		From:    "Alice",
		Message: "let's grab a :beer:",
		To:      "+44782669123",
	})
	fmt.Println(rsp, err)
}
