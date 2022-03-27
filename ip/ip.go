package ip

import (
	"go.m3o.com/client"
)

type Ip interface {
	Lookup(*LookupRequest) (*LookupResponse, error)
}

func NewIpService(token string) *IpService {
	return &IpService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type IpService struct {
	client *client.Client
}

// Lookup the geolocation information for an IP address
func (t *IpService) Lookup(request *LookupRequest) (*LookupResponse, error) {

	rsp := &LookupResponse{}
	return rsp, t.client.Call("ip", "Lookup", request, rsp)

}

type LookupRequest struct {
	// IP to lookup
	Ip string `json:"ip,omitempty"`
}

type LookupResponse struct {
	// Autonomous system number
	Asn int32 `json:"asn,omitempty"`
	// Name of the city
	City string `json:"city,omitempty"`
	// Name of the continent
	Continent string `json:"continent,omitempty"`
	// Name of the country
	Country string `json:"country,omitempty"`
	// IP of the query
	Ip string `json:"ip,omitempty"`
	// Latitude e.g 52.523219
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude e.g 13.428555
	Longitude float64 `json:"longitude,omitempty"`
	// Timezone e.g Europe/Rome
	Timezone string `json:"timezone,omitempty"`
}
