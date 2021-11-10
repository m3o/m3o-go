# Routing

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Routing/api](https://m3o.com/Routing/api).

Endpoints:

## Directions

Turn by turn directions from a start point to an end point including maneuvers and bearings


[https://m3o.com/routing/api#Directions](https://m3o.com/routing/api#Directions)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/routing"
)

// Turn by turn directions from a start point to an end point including maneuvers and bearings
func TurnByTurnDirections() {
	routingService := routing.NewRoutingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := routingService.Directions(&routing.DirectionsRequest{
		Destination: &routing.Point{
	Latitude: 52.529407,
	Longitude: 13.397634,
},
Origin: &routing.Point{
	Latitude: 52.517037,
	Longitude: 13.38886,
},

	})
	fmt.Println(rsp, err)
	
}
```
## Route

Retrieve a route as a simple list of gps points along with total distance and estimated duration


[https://m3o.com/routing/api#Route](https://m3o.com/routing/api#Route)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/routing"
)

// Retrieve a route as a simple list of gps points along with total distance and estimated duration
func GpsPointsForAroute() {
	routingService := routing.NewRoutingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := routingService.Route(&routing.RouteRequest{
		Destination: &routing.Point{
	Latitude: 52.529407,
	Longitude: 13.397634,
},
Origin: &routing.Point{
	Latitude: 52.517037,
	Longitude: 13.38886,
},

	})
	fmt.Println(rsp, err)
	
}
```
## Eta

Get the eta for a route from origin to destination. The eta is an estimated time based on car routes


[https://m3o.com/routing/api#Eta](https://m3o.com/routing/api#Eta)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/routing"
)

// Get the eta for a route from origin to destination. The eta is an estimated time based on car routes
func EtaFromPointAtoPointB() {
	routingService := routing.NewRoutingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := routingService.Eta(&routing.EtaRequest{
		Destination: &routing.Point{
	Latitude: 52.529407,
	Longitude: 13.397634,
},
Origin: &routing.Point{
	Latitude: 52.517037,
	Longitude: 13.38886,
},

	})
	fmt.Println(rsp, err)
	
}
```
