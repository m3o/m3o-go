package ai

import (
	"go.m3o.com/client"
)

type Ai interface {
	Chat(*ChatRequest) (*ChatResponse, error)
	Complete(*CompleteRequest) (*CompleteResponse, error)
	Edit(*EditRequest) (*EditResponse, error)
	Generate(*GenerateRequest) (*GenerateResponse, error)
	Moderate(*ModerateRequest) (*ModerateResponse, error)
	Stream(*StreamRequest) (*StreamResponseStream, error)
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

// Make a request to ChatGPT
func (t *AiService) Chat(request *ChatRequest) (*ChatResponse, error) {

	rsp := &ChatResponse{}
	return rsp, t.client.Call("ai", "Chat", request, rsp)

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

// Generate an image from prompt
func (t *AiService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("ai", "Generate", request, rsp)

}

// Moderate hate speech
func (t *AiService) Moderate(request *ModerateRequest) (*ModerateResponse, error) {

	rsp := &ModerateResponse{}
	return rsp, t.client.Call("ai", "Moderate", request, rsp)

}

// Stream a response from chatgpt
func (t *AiService) Stream(request *StreamRequest) (*StreamResponseStream, error) {
	stream, err := t.client.Stream("ai", "Stream", request)
	if err != nil {
		return nil, err
	}
	return &StreamResponseStream{
		stream: stream,
	}, nil

}

type StreamResponseStream struct {
	stream *client.Stream
}

func (t *StreamResponseStream) Recv() (*StreamResponse, error) {
	var rsp StreamResponse
	if err := t.stream.Recv(&rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

type ChatRequest struct {
	// context for the call
	Context []Context `json:"context,omitempty"`
	// the model e.g gpt-3.5-turbo-16k
	Model string `json:"model,omitempty"`
	// the prompt
	Prompt string `json:"prompt,omitempty"`
	// role e.g system or user
	Role string `json:"role,omitempty"`
}

type ChatResponse struct {
	// the response
	Reply string `json:"reply,omitempty"`
}

type CompleteRequest struct {
	// input to pass in
	Text string `json:"text,omitempty"`
}

type CompleteResponse struct {
	// text returned
	Text string `json:"text,omitempty"`
}

type Context struct {
	// prompt used
	Prompt string `json:"prompt,omitempty"`
	// response for prompt
	Reply string `json:"reply,omitempty"`
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

type GenerateRequest struct {
	// number of images to generate (max 10)
	Limit int32 `json:"limit,omitempty"`
	// size of image 256x256, 512x512, 1024x1024
	Size string `json:"size,omitempty"`
	// text description of image
	Text string `json:"text,omitempty"`
}

type GenerateResponse struct {
	// image urls
	Images []Image `json:"images,omitempty"`
}

type Image struct {
	// base64 encoded
	Base64 string `json:"base64,omitempty"`
	// image url
	Url string `json:"url,omitempty"`
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

type StreamRequest struct {
	// the potential model e.g gpt-4
	Model string `json:"model,omitempty"`
	// the prompt to provide
	Prompt string `json:"prompt,omitempty"`
}

type StreamResponse struct {
	// whether its a complete or partial response
	Partial bool `json:"partial,omitempty"`
	// a set of words in the response
	Words []string `json:"words,omitempty"`
}
