package helloworld

import (
	"go.m3o.com/client"
)

type Helloworld interface {
	Call(*CallRequest) (*CallResponse, error)
	Stream(*StreamRequest) (*StreamResponseStream, error)
}

func NewHelloworldService(token string) *HelloworldService {
	return &HelloworldService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type HelloworldService struct {
	client *client.Client
}

// Return a personalised Hello message
func (t *HelloworldService) Call(request *CallRequest) (*CallResponse, error) {

	rsp := &CallResponse{}
	return rsp, t.client.Call("helloworld", "Call", request, rsp)

}

// Stream a personalised Hello message
func (t *HelloworldService) Stream(request *StreamRequest) (*StreamResponseStream, error) {
	stream, err := t.client.Stream("helloworld", "Stream", request)
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

type CallRequest struct {
	// name to append to the message e.g Alice
	Name string `json:"name,omitempty"`
}

type CallResponse struct {
	// message including name e.g Hello Alice
	Message string `json:"message,omitempty"`
}

type StreamRequest struct {
	// the number of messages to send back
	Messages int64 `json:"messages,string,omitempty"`
	// name to append to the message e.g Alice
	Name string `json:"name,omitempty"`
}

type StreamResponse struct {
	// message including name e.g Hello Alice
	Message string `json:"message,omitempty"`
}
