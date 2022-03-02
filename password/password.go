package password

import (
	"go.m3o.com/client"
)

type Password interface {
	Generate(*GenerateRequest) (*GenerateResponse, error)
}

func NewPasswordService(token string) *PasswordService {
	return &PasswordService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PasswordService struct {
	client *client.Client
}

// Generate a strong random password
func (t *PasswordService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("password", "Generate", request, rsp)

}

type GenerateRequest struct {
	// password length; defaults to 16 chars
	Length int32 `json:"length"`
}

type GenerateResponse struct {
	// The generated password
	Password string `json:"password"`
}
