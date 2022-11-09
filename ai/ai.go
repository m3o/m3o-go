package ai

import (
	"go.m3o.com/client"
)

type Ai interface {
	Call(*CallRequest) (*CallResponse, error)
	Check(*CheckRequest) (*CheckResponse, error)
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
func (t *AiService) Call(request *CallRequest) (*CallResponse, error) {

	rsp := &CallResponse{}
	return rsp, t.client.Call("ai", "Call", request, rsp)

}

// Check or edit text/code
func (t *AiService) Check(request *CheckRequest) (*CheckResponse, error) {

	rsp := &CheckResponse{}
	return rsp, t.client.Call("ai", "Check", request, rsp)

}

// Moderate hate speech
func (t *AiService) Moderate(request *ModerateRequest) (*ModerateResponse, error) {

	rsp := &ModerateResponse{}
	return rsp, t.client.Call("ai", "Moderate", request, rsp)

}

type CallRequest struct {
	// text to pass in
	Text string `json:"text,omitempty"`
}

type CallResponse struct {
	// text returned
	Text string `json:"text,omitempty"`
}

type CheckRequest struct {
	// instruction hint e.g check the grammar
	Instruction string `json:"instruction,omitempty"`
	// text/code to check
	Text string `json:"text,omitempty"`
}

type CheckResponse struct {
	// response output
	Text string `json:"text,omitempty"`
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