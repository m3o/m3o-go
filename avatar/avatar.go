package avatar

import (
	"go.m3o.com/client"
)

type Avatar interface {
	Generate(*GenerateRequest) (*GenerateResponse, error)
}

func NewAvatarService(token string) *AvatarService {
	return &AvatarService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AvatarService struct {
	client *client.Client
}

// Generate an unique avatar
func (t *AvatarService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("avatar", "Generate", request, rsp)

}

type GenerateRequest struct {
	// encode format of avatar image: `png` or `jpeg`; default is `jpeg`
	Format string `json:"format,omitempty"`
	// avatar's gender: `male` or `female`; default is `male`
	Gender string `json:"gender,omitempty"`
	// set to true to upload to the M3O CDN and receive the url
	Upload bool `json:"upload,omitempty"`
	// avatar's username, unique username will generate the unique avatar;
	// if empty, every request generates a random avatar;
	// if upload == true, username will be the CDN filename rather than a random uuid string
	Username string `json:"username,omitempty"`
}

type GenerateResponse struct {
	// base64 encoded string of the avatar image
	Base64 string `json:"base64,omitempty"`
	// M3O's CDN url of the avatar image
	Url string `json:"url,omitempty"`
}
