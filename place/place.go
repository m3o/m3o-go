package place

import (
	"go.m3o.com/client"
)

type Place interface {
	Nearby(*NearbyRequest) (*NearbyResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewPlaceService(token string) *PlaceService {
	return &PlaceService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PlaceService struct {
	client *client.Client
}

// Find places nearby using a location
func (t *PlaceService) Nearby(request *NearbyRequest) (*NearbyResponse, error) {

	rsp := &NearbyResponse{}
	return rsp, t.client.Call("place", "Nearby", request, rsp)

}

// Search for places by text query
func (t *PlaceService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("place", "Search", request, rsp)

}

type AutocompleteRequest struct {
}

type AutocompleteResponse struct {
}

type NearbyRequest struct {
	// Keyword to include in the search
	Keyword string `json:"keyword,omitempty"`
	// specify the location by lat,lng e.g -33.8670522,-151.1957362
	Location string `json:"location,omitempty"`
	// Name of the place to search for
	Name string `json:"name,omitempty"`
	// Whether the place is open now
	OpenNow bool `json:"open_now,omitempty"`
	// radius in meters within which to search
	Radius int32 `json:"radius,omitempty"`
	// Type of place. https://developers.google.com/maps/documentation/places/web-service/supported_types
	Type string `json:"type,omitempty"`
}

type NearbyResponse struct {
	Results []Result `json:"results,omitempty"`
}

type Result struct {
	// address of place
	Address string `json:"address,omitempty"`
	// url of an icon
	IconUrl string `json:"icon_url,omitempty"`
	// lat/lng of place
	Location string `json:"location,omitempty"`
	// name of the place
	Name string `json:"name,omitempty"`
	// open now
	OpenNow bool `json:"open_now,omitempty"`
	// opening hours
	OpeningHours string `json:"opening_hours,omitempty"`
	// rating from 1.0 to 5.0
	Rating float64 `json:"rating,omitempty"`
	// type of location
	Type string `json:"type,omitempty"`
	// feature types
	Types []string `json:"types,omitempty"`
	// simplified address
	Vicinity string `json:"vicinity,omitempty"`
}

type SearchRequest struct {
	// the location by lat,lng e.g -33.8670522,-151.1957362
	Location string `json:"location,omitempty"`
	// Whether the place is open now
	OpenNow bool `json:"open_now,omitempty"`
	// the text string on which to search, for example: "restaurant"
	Query string `json:"query,omitempty"`
	// radius in meters within which to search
	Radius int32 `json:"radius,omitempty"`
	// Type of place. https://developers.google.com/maps/documentation/places/web-service/supported_types
	Type string `json:"type,omitempty"`
}

type SearchResponse struct {
	Results []Result `json:"results,omitempty"`
}
