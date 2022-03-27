package analytics

import (
	"go.m3o.com/client"
)

type Analytics interface {
	Delete(*DeleteRequest) (*DeleteResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Track(*TrackRequest) (*TrackResponse, error)
}

func NewAnalyticsService(token string) *AnalyticsService {
	return &AnalyticsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AnalyticsService struct {
	client *client.Client
}

// Delete an event
func (t *AnalyticsService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("analytics", "Delete", request, rsp)

}

// List all events
func (t *AnalyticsService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("analytics", "List", request, rsp)

}

// Get a single event
func (t *AnalyticsService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("analytics", "Read", request, rsp)

}

// Track an event, it will be created if it doesn't exist
func (t *AnalyticsService) Track(request *TrackRequest) (*TrackResponse, error) {

	rsp := &TrackResponse{}
	return rsp, t.client.Call("analytics", "Track", request, rsp)

}

type DeleteRequest struct {
	Name string `json:"name,omitempty"`
}

type DeleteResponse struct {
	Event *Event `json:"event,omitempty"`
}

type Event struct {
	// time at which the event was created
	Created string `json:"created,omitempty"`
	// event name
	Name string `json:"name,omitempty"`
	// the amount of times the event was triggered
	Value int64 `json:"value,string,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	Events []Event `json:"events,omitempty"`
}

type ReadRequest struct {
	Name string `json:"name,omitempty"`
}

type ReadResponse struct {
	Event *Event `json:"event,omitempty"`
}

type TrackRequest struct {
	// event name
	Name string `json:"name,omitempty"`
}

type TrackResponse struct {
}
