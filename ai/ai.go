package ai

import (
	"go.m3o.com/client"
)

type Ai interface {
	Complete(*CompleteRequest) (*CompleteResponse, error)
	Edit(*EditRequest) (*EditResponse, error)
	Image(*ImageRequest) (*ImageResponse, error)
	Moderate(*ModerateRequest) (*ModerateResponse, error)
}

func NewAiService(token string) *AiService {
	return &AiService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AiService struct {
	client *client.Client
}

// Make a request to the AI
func (t *AiService) Complete(request *CompleteRequest) (*CompleteResponse, error) {

	rsp := &CompleteResponse{}
	return rsp, t.client.Call("ai", "Complete", request, rsp)

}

// Edit or edit prompt/code
func (t *AiService) Edit(request *EditRequest) (*EditResponse, error) {

	rsp := &EditResponse{}
	return rsp, t.client.Call("ai", "Edit", request, rsp)

}

// Generage an image from prompt
func (t *AiService) Image(request *ImageRequest) (*ImageResponse, error) {

	rsp := &ImageResponse{}
	return rsp, t.client.Call("ai", "Image", request, rsp)

}

// Moderate hate speech
func (t *AiService) Moderate(request *ModerateRequest) (*ModerateResponse, error) {

	rsp := &ModerateResponse{}
	return rsp, t.client.Call("ai", "Moderate", request, rsp)

}

type CompleteRequest struct {
	// input to pass in
	Text string `json:"text,omitempty"`
}

type CompleteResponse struct {
	// text returned
	Text string `json:"text,omitempty"`
}

type EditRequest struct {
	// instruction hint e.g check the grammar
	Instruction string `json:"instruction,omitempty"`
	// text/code to check
	Text string `json:"text,omitempty"`
}

type EditResponse struct {
	// response output
	Text string `json:"text,omitempty"`
}

type Image struct {
	// base64 encoded
	Base64 string `json:"base64,omitempty"`
	// image url
	Url string `json:"url,omitempty"`
}

type ImageRequest struct {
	// number of images to generate (max 10)
	Limit int32 `json:"limit,omitempty"`
	// size of image 256x256, 512x512, 1024x1024
	Size string `json:"size,omitempty"`
	// text description of image
	Text string `json:"text,omitempty"`
}

type ImageResponse struct {
	// image urls
	Images []Image `json:"images,omitempty"`
}

type ModerateRequest struct {
	// text to check
	Text string `json:"text,omitempty"`
}

type ModerateResponse struct {
	// categories tested and identified
	Categories map[string]bool `json:"categories,omitempty"`
	// whether it was flagged or not
	Flagged bool `json:"flagged,omitempty"`
	// related scores
	Scores map[string]float64 `json:"scores,omitempty"`
}
