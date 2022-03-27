package geocoding

import (
	"go.m3o.com/client"
)

type Geocoding interface {
	Lookup(*LookupRequest) (*LookupResponse, error)
	Reverse(*ReverseRequest) (*ReverseResponse, error)
}

func NewGeocodingService(token string) *GeocodingService {
	return &GeocodingService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type GeocodingService struct {
	client *client.Client
}

// Lookup returns a geocoded address including normalized address and gps coordinates. All fields are optional, provide more to get more accurate results
func (t *GeocodingService) Lookup(request *LookupRequest) (*LookupResponse, error) {

	rsp := &LookupResponse{}
	return rsp, t.client.Call("geocoding", "Lookup", request, rsp)

}

// Reverse lookup an address from gps coordinates
func (t *GeocodingService) Reverse(request *ReverseRequest) (*ReverseResponse, error) {

	rsp := &ReverseResponse{}
	return rsp, t.client.Call("geocoding", "Reverse", request, rsp)

}

type Address struct {
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
	LineOne  string `json:"line_one,omitempty"`
	LineTwo  string `json:"line_two,omitempty"`
	Postcode string `json:"postcode,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type LookupRequest struct {
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
	Postcode string `json:"postcode,omitempty"`
}

type LookupResponse struct {
	Address  *Address  `json:"address,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type ReverseRequest struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type ReverseResponse struct {
	Address  *Address  `json:"address,omitempty"`
	Location *Location `json:"location,omitempty"`
}
