package time

import (
	"go.m3o.com/client"
)

type Time interface {
	Now(*NowRequest) (*NowResponse, error)
	Zone(*ZoneRequest) (*ZoneResponse, error)
}

func NewTimeService(token string) *TimeService {
	return &TimeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type TimeService struct {
	client *client.Client
}

// Get the current time
func (t *TimeService) Now(request *NowRequest) (*NowResponse, error) {

	rsp := &NowResponse{}
	return rsp, t.client.Call("time", "Now", request, rsp)

}

// Get the timezone info for a specific location
func (t *TimeService) Zone(request *ZoneRequest) (*ZoneResponse, error) {

	rsp := &ZoneResponse{}
	return rsp, t.client.Call("time", "Zone", request, rsp)

}

type NowRequest struct {
	// optional location, otherwise returns UTC
	Location string `json:"location,omitempty"`
}

type NowResponse struct {
	// the current time as HH:MM:SS
	Localtime string `json:"localtime,omitempty"`
	// the location as Europe/London
	Location string `json:"location,omitempty"`
	// timestamp as 2006-01-02T15:04:05.999999999Z07:00
	Timestamp string `json:"timestamp,omitempty"`
	// the timezone as BST
	Timezone string `json:"timezone,omitempty"`
	// the unix timestamp
	Unix int64 `json:"unix,string,omitempty"`
}

type ZoneRequest struct {
	// location to lookup e.g postcode, city, ip address
	Location string `json:"location,omitempty"`
}

type ZoneResponse struct {
	// the abbreviated code e.g BST
	Abbreviation string `json:"abbreviation,omitempty"`
	// country of the timezone
	Country string `json:"country,omitempty"`
	// is daylight savings
	Dst bool `json:"dst,omitempty"`
	// e.g 51.42
	Latitude float64 `json:"latitude,omitempty"`
	// the local time
	Localtime string `json:"localtime,omitempty"`
	// location requested
	Location string `json:"location,omitempty"`
	// e.g -0.37
	Longitude float64 `json:"longitude,omitempty"`
	// UTC offset in hours
	Offset int32 `json:"offset,omitempty"`
	// region of timezone
	Region string `json:"region,omitempty"`
	// the timezone e.g Europe/London
	Timezone string `json:"timezone,omitempty"`
}
