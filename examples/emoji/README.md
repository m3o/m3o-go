# Emoji

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/emoji/api](https://m3o.com/emoji/api).

Endpoints:

## Find

Find an emoji by its alias e.g :beer:


[https://m3o.com/emoji/api#Find](https://m3o.com/emoji/api#Find)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Find an emoji by its alias e.g :beer:
func FindEmoji() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Find(&emoji.FindRequest{
		Alias: ":beer:",

	})
	fmt.Println(rsp, err)
	
}
```
## Flag

Get the flag for a country. Requires country code e.g GB for great britain


[https://m3o.com/emoji/api#Flag](https://m3o.com/emoji/api#Flag)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Get the flag for a country. Requires country code e.g GB for great britain
func GetFlagByCountryCode() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Flag(&emoji.FlagRequest{
		Code: "GB",

	})
	fmt.Println(rsp, err)
	
}
```
## Print

Print text and renders the emojis with aliases e.g
let's grab a :beer: becomes let's grab a 🍺


[https://m3o.com/emoji/api#Print](https://m3o.com/emoji/api#Print)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/emoji"
)

// Print text and renders the emojis with aliases e.g
// let's grab a :beer: becomes let's grab a 🍺
func PrintTextIncludingEmoji() {
	emojiService := emoji.NewEmojiService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := emojiService.Print(&emoji.PrintRequest{
		Text: "let's grab a :beer:",

	})
	fmt.Println(rsp, err)
	
}
```
