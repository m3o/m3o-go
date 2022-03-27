package joke

import (
	"go.m3o.com/client"
)

type Joke interface {
	Random(*RandomRequest) (*RandomResponse, error)
}

func NewJokeService(token string) *JokeService {
	return &JokeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type JokeService struct {
	client *client.Client
}

// Get a random joke
func (t *JokeService) Random(request *RandomRequest) (*RandomResponse, error) {

	rsp := &RandomResponse{}
	return rsp, t.client.Call("joke", "Random", request, rsp)

}

type JokeInfo struct {
	Body     string `json:"body,omitempty"`
	Category string `json:"category,omitempty"`
	Id       string `json:"id,omitempty"`
	Source   string `json:"source,omitempty"`
	Title    string `json:"title,omitempty"`
}

type RandomRequest struct {
	// the count of random jokes want, maximum: 10
	Count int32 `json:"count,omitempty"`
}

type RandomResponse struct {
	Jokes []JokeInfo `json:"jokes,omitempty"`
}
