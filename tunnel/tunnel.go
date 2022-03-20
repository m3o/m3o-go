package tunnel

import (
	"go.m3o.com/client"
)

type Tunnel interface {
	Send(*SendRequest) (*SendResponse, error)
}

func NewTunnelService(token string) *TunnelService {
	return &TunnelService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type TunnelService struct {
	client *client.Client
}

// Send a request through the tunnel
func (t *TunnelService) Send(request *SendRequest) (*SendResponse, error) {

	rsp := &SendResponse{}
	return rsp, t.client.Call("tunnel", "Send", request, rsp)

}

type SendRequest struct {
	// body of the request
	Body string `json:"body"`
	// headers to include e.g Content-Type: application/json
	Headers map[string]string `json:"headers"`
	// host to send to e.g www.google.com
	Host string `json:"host"`
	// method of the request e.g GET, POST, DELETE
	Method string `json:"method"`
	// request params to include
	Params map[string]string `json:"params"`
	// path to request e.g /news
	Path string `json:"path"`
	// full url e.g https://www.google.com/news
	Url string `json:"url"`
}

type SendResponse struct {
	// body of the response
	Body string `json:"body"`
	// headers to include e.g Content-Type: application/json
	Headers map[string]string `json:"headers"`
	// the status
	Status string `json:"status"`
	// the status code
	StatusCode int32 `json:"status_code"`
}
