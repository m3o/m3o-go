package example

import (
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Get the flag for a country. Requires country code e.g GB for great britain
func GetFlagByCountryCode() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Flag(&emoji.FlagRequest{})
	fmt.Println(rsp, err)
}
