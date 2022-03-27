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

// Generate a strong random password. Use the switches to control which character types are included, defaults to using all of them
func (t *PasswordService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("password", "Generate", request, rsp)

}

type GenerateRequest struct {
	// password length; defaults to 8 chars
	Length int32 `json:"length,omitempty"`
	// include lowercase letters
	Lowercase bool `json:"lowercase,omitempty"`
	// include numbers
	Numbers bool `json:"numbers,omitempty"`
	// include special characters (!@#$%&*)
	Special bool `json:"special,omitempty"`
	// include uppercase letters
	Uppercase bool `json:"uppercase,omitempty"`
}

type GenerateResponse struct {
	// The generated password
	Password string `json:"password,omitempty"`
}
