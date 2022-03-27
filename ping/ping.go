package ping

import (
	"go.m3o.com/client"
)

type Ping interface {
	Ip(*IpRequest) (*IpResponse, error)
	Tcp(*TcpRequest) (*TcpResponse, error)
	Url(*UrlRequest) (*UrlResponse, error)
}

func NewPingService(token string) *PingService {
	return &PingService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PingService struct {
	client *client.Client
}

// Ping an IP address
func (t *PingService) Ip(request *IpRequest) (*IpResponse, error) {

	rsp := &IpResponse{}
	return rsp, t.client.Call("ping", "Ip", request, rsp)

}

// Ping a TCP port is open
func (t *PingService) Tcp(request *TcpRequest) (*TcpResponse, error) {

	rsp := &TcpResponse{}
	return rsp, t.client.Call("ping", "Tcp", request, rsp)

}

// Ping a HTTP URL
func (t *PingService) Url(request *UrlRequest) (*UrlResponse, error) {

	rsp := &UrlResponse{}
	return rsp, t.client.Call("ping", "Url", request, rsp)

}

type IpRequest struct {
	// address to ping
	Address string `json:"address,omitempty"`
}

type IpResponse struct {
	// average latency e.g 10ms
	Latency string `json:"latency,omitempty"`
	// response status
	Status string `json:"status,omitempty"`
}

type TcpRequest struct {
	// address to dial
	Address string `json:"address,omitempty"`
	// optional data to send
	Data string `json:"data,omitempty"`
}

type TcpResponse struct {
	// response data if any
	Data string `json:"data,omitempty"`
	// response status
	Status string `json:"status,omitempty"`
}

type UrlRequest struct {
	// address to use
	Address string `json:"address,omitempty"`
	// method of the call
	Method string `json:"method,omitempty"`
}

type UrlResponse struct {
	// the response code
	Code int32 `json:"code,omitempty"`
	// the response status
	Status string `json:"status,omitempty"`
}
