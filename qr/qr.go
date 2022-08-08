package qr

import (
	"go.m3o.com/client"
)

type Qr interface {
	Codes(*CodesRequest) (*CodesResponse, error)
	Generate(*GenerateRequest) (*GenerateResponse, error)
}

func NewQrService(token string) *QrService {
	return &QrService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type QrService struct {
	client *client.Client
}

// List your QR codes
func (t *QrService) Codes(request *CodesRequest) (*CodesResponse, error) {

	rsp := &CodesResponse{}
	return rsp, t.client.Call("qr", "Codes", request, rsp)

}

// Generate a QR code with a specific text and size
func (t *QrService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("qr", "Generate", request, rsp)

}

type Code struct {
	// time of creation
	Created string `json:"created,omitempty"`
	// file name
	File string `json:"file,omitempty"`
	Id   string `json:"id,omitempty"`
	// text of the QR code
	Text string `json:"text,omitempty"`
}

type CodesRequest struct {
}

type CodesResponse struct {
	Codes []Code `json:"codes,omitempty"`
}

type GenerateRequest struct {
	// the size (height and width) in pixels of the generated QR code. Defaults to 256
	Size int64 `json:"size,string,omitempty"`
	// the text to encode as a QR code (URL, phone number, email, etc)
	Text string `json:"text,omitempty"`
}

type GenerateResponse struct {
	// link to the QR code image in PNG format
	Qr string `json:"qr,omitempty"`
}
