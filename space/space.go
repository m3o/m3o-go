package space

import (
	"go.m3o.com/client"
)

func NewSpaceService(token string) *SpaceService {
	return &SpaceService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SpaceService struct {
	client *client.Client
}

// Vote to have the Space api launched faster!
func (t *SpaceService) Vote(request *VoteRequest) (*VoteResponse, error) {

	rsp := &VoteResponse{}
	return rsp, t.client.Call("space", "Vote", request, rsp)

}

type VoteRequest struct {
	// optional message
	Message string `json:"message"`
}

type VoteResponse struct {
	// response message
	Message string `json:"message"`
}
