package emoji

import (
	"go.m3o.com/client"
)

type Emoji interface {
	Find(*FindRequest) (*FindResponse, error)
	Flag(*FlagRequest) (*FlagResponse, error)
	Print(*PrintRequest) (*PrintResponse, error)
}

func NewEmojiService(token string) *EmojiService {
	return &EmojiService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EmojiService struct {
	client *client.Client
}

// Find an emoji by its alias e.g :beer:
func (t *EmojiService) Find(request *FindRequest) (*FindResponse, error) {

	rsp := &FindResponse{}
	return rsp, t.client.Call("emoji", "Find", request, rsp)

}

// Get the flag for a country. Requires country code e.g GB for great britain
func (t *EmojiService) Flag(request *FlagRequest) (*FlagResponse, error) {

	rsp := &FlagResponse{}
	return rsp, t.client.Call("emoji", "Flag", request, rsp)

}

// Print text and renders the emojis with aliases e.g
// let's grab a :beer: becomes let's grab a 🍺
func (t *EmojiService) Print(request *PrintRequest) (*PrintResponse, error) {

	rsp := &PrintResponse{}
	return rsp, t.client.Call("emoji", "Print", request, rsp)

}

type FindRequest struct {
	// the alias code e.g :beer:
	Alias string `json:"alias,omitempty"`
}

type FindResponse struct {
	// the unicode emoji 🍺
	Emoji string `json:"emoji,omitempty"`
}

type FlagRequest struct {
	// country code e.g GB
	Code string `json:"code,omitempty"`
}

type FlagResponse struct {
	// the emoji flag
	Flag string `json:"flag,omitempty"`
}

type PrintRequest struct {
	// text including any alias e.g let's grab a :beer:
	Text string `json:"text,omitempty"`
}

type PrintResponse struct {
	// text with rendered emojis
	Text string `json:"text,omitempty"`
}
