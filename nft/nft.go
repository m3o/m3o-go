package nft

import (
	"go.m3o.com/client"
)

func NewNftService(token string) *NftService {
	return &NftService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type NftService struct {
	client *client.Client
}

// Vote to have the NFT api launched faster!
func (t *NftService) Vote(request *VoteRequest) (*VoteResponse, error) {

	rsp := &VoteResponse{}
	return rsp, t.client.Call("nft", "Vote", request, rsp)

}

type VoteRequest struct {
	// optional message
	Message string `json:"message"`
}

type VoteResponse struct {
	// response message
	Message string `json:"message"`
}
