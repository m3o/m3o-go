package location

import (
	"go.m3o.com/client"
)

type Location interface {
	Read(*ReadRequest) (*ReadResponse, error)
	Save(*SaveRequest) (*SaveResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewLocationService(token string) *LocationService {
	return &LocationService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type LocationService struct {
	client *client.Client
}

// Read an entity by its ID
func (t *LocationService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("location", "Read", request, rsp)

}

// Save an entity's current position
func (t *LocationService) Save(request *SaveRequest) (*SaveResponse, error) {

	rsp := &SaveResponse{}
	return rsp, t.client.Call("location", "Save", request, rsp)

}

// Search for entities in a given radius
func (t *LocationService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("location", "Search", request, rsp)

}

type Entity struct {
	Id       string `json:"id,omitempty"`
	Location *Point `json:"location,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Point struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Timestamp int64   `json:"timestamp,string,omitempty"`
}

type ReadRequest struct {
	// the entity id
	Id string `json:"id,omitempty"`
}

type ReadResponse struct {
	Entity *Entity `json:"entity,omitempty"`
}

type SaveRequest struct {
	Entity *Entity `json:"entity,omitempty"`
}

type SaveResponse struct {
}

type SearchRequest struct {
	// Central position to search from
	Center *Point `json:"center,omitempty"`
	// Maximum number of entities to return
	NumEntities int64 `json:"numEntities,string,omitempty"`
	// radius in meters
	Radius float64 `json:"radius,omitempty"`
	// type of entities to filter
	Type string `json:"type,omitempty"`
}

type SearchResponse struct {
	Entities []Entity `json:"entities,omitempty"`
}
