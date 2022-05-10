package wordle

import (
	"go.m3o.com/client"
)

type Wordle interface {
	Guess(*GuessRequest) (*GuessResponse, error)
	Next(*NextRequest) (*NextResponse, error)
}

func NewWordleService(token string) *WordleService {
	return &WordleService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type WordleService struct {
	client *client.Client
}

// Make a guess
func (t *WordleService) Guess(request *GuessRequest) (*GuessResponse, error) {

	rsp := &GuessResponse{}
	return rsp, t.client.Call("wordle", "Guess", request, rsp)

}

// When does the next game start
func (t *WordleService) Next(request *NextRequest) (*NextResponse, error) {

	rsp := &NextResponse{}
	return rsp, t.client.Call("wordle", "Next", request, rsp)

}

type Char struct {
	// whether it was correct
	Correct bool `json:"correct,omitempty"`
	// whether it's in the word
	InWord bool `json:"in_word,omitempty"`
	// the character itself
	Letter string `json:"letter,omitempty"`
	// position in the string
	Position int32 `json:"position,omitempty"`
}

type Guess struct {
	// individual characters
	Chars []Char `json:"chars,omitempty"`
	// the highlighted word e.g n[o]is{e}
	// where [ ] is correct, { } is in word
	Highlight string `json:"highlight,omitempty"`
	// the full guess word
	Word string `json:"word,omitempty"`
}

type GuessRequest struct {
	// player
	Player string `json:"player,omitempty"`
	// guess word
	Word string `json:"word,omitempty"`
}

type GuessResponse struct {
	// the actual word if failed
	Answer string `json:"answer,omitempty"`
	// whether it was correct
	Correct bool `json:"correct,omitempty"`
	// the guess words tried
	Guesses []Guess `json:"guesses,omitempty"`
	// informational message
	Status string `json:"status,omitempty"`
	// number of tries left
	TriesLeft int32 `json:"tries_left,omitempty"`
}

type NextRequest struct {
}

type NextResponse struct {
	// in hh:mm:ss
	Duration string `json:"duration,omitempty"`
	// number of seconds
	Seconds int32 `json:"seconds,omitempty"`
}
