package routing

import (
	"go.m3o.com/client"
)

type Routing interface {
	Directions(*DirectionsRequest) (*DirectionsResponse, error)
	Eta(*EtaRequest) (*EtaResponse, error)
	Route(*RouteRequest) (*RouteResponse, error)
}

func NewRoutingService(token string) *RoutingService {
	return &RoutingService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type RoutingService struct {
	client *client.Client
}

// Turn by turn directions from a start point to an end point including maneuvers and bearings
func (t *RoutingService) Directions(request *DirectionsRequest) (*DirectionsResponse, error) {

	rsp := &DirectionsResponse{}
	return rsp, t.client.Call("routing", "Directions", request, rsp)

}

// Get the eta for a route from origin to destination. The eta is an estimated time based on car routes
func (t *RoutingService) Eta(request *EtaRequest) (*EtaResponse, error) {

	rsp := &EtaResponse{}
	return rsp, t.client.Call("routing", "Eta", request, rsp)

}

// Retrieve a route as a simple list of gps points along with total distance and estimated duration
func (t *RoutingService) Route(request *RouteRequest) (*RouteResponse, error) {

	rsp := &RouteResponse{}
	return rsp, t.client.Call("routing", "Route", request, rsp)

}

type Direction struct {
	// distance to travel in meters
	Distance float64 `json:"distance,omitempty"`
	// duration to travel in seconds
	Duration float64 `json:"duration,omitempty"`
	// human readable instruction
	Instruction string `json:"instruction,omitempty"`
	// intersections on route
	Intersections []Intersection `json:"intersections,omitempty"`
	// maneuver to take
	Maneuver *Maneuver `json:"maneuver,omitempty"`
	// street name or location
	Name string `json:"name,omitempty"`
	// alternative reference
	Reference string `json:"reference,omitempty"`
}

type DirectionsRequest struct {
	// The destination of the journey
	Destination *Point `json:"destination,omitempty"`
	// The staring point for the journey
	Origin *Point `json:"origin,omitempty"`
}

type DirectionsResponse struct {
	// Turn by turn directions
	Directions []Direction `json:"directions,omitempty"`
	// Estimated distance of the route in meters
	Distance float64 `json:"distance,omitempty"`
	// Estimated duration of the route in seconds
	Duration float64 `json:"duration,omitempty"`
	// The waypoints on the route
	Waypoints []Waypoint `json:"waypoints,omitempty"`
}

type EtaRequest struct {
	// The end point for the eta calculation
	Destination *Point `json:"destination,omitempty"`
	// The starting point for the eta calculation
	Origin *Point `json:"origin,omitempty"`
	// speed in kilometers
	Speed float64 `json:"speed,omitempty"`
	// type of transport. Only "car" is supported currently.
	Type string `json:"type,omitempty"`
}

type EtaResponse struct {
	// eta in seconds
	Duration float64 `json:"duration,omitempty"`
}

type Intersection struct {
	Bearings []float64 `json:"bearings,omitempty"`
	Location *Point    `json:"location,omitempty"`
}

type Maneuver struct {
	Action        string  `json:"action,omitempty"`
	BearingAfter  float64 `json:"bearing_after,omitempty"`
	BearingBefore float64 `json:"bearing_before,omitempty"`
	Direction     string  `json:"direction,omitempty"`
	Location      *Point  `json:"location,omitempty"`
}

type Point struct {
	// Lat e.g 52.523219
	Latitude float64 `json:"latitude,omitempty"`
	// Long e.g 13.428555
	Longitude float64 `json:"longitude,omitempty"`
}

type RouteRequest struct {
	// Point of destination for the trip
	Destination *Point `json:"destination,omitempty"`
	// Point of origin for the trip
	Origin *Point `json:"origin,omitempty"`
}

type RouteResponse struct {
	// estimated distance in meters
	Distance float64 `json:"distance,omitempty"`
	// estimated duration in seconds
	Duration float64 `json:"duration,omitempty"`
	// waypoints on the route
	Waypoints []Waypoint `json:"waypoints,omitempty"`
}

type Waypoint struct {
	// gps point coordinates
	Location *Point `json:"location,omitempty"`
	// street name or related reference
	Name string `json:"name,omitempty"`
}
