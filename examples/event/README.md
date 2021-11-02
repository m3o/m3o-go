# Event

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Event/api](https://m3o.com/Event/api).

Endpoints:

## Publish

Publish a message to the event. Specify a topic to group messages for a specific topic.


[https://m3o.com/event/api#Publish](https://m3o.com/event/api#Publish)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Publish a message to the event. Specify a topic to group messages for a specific topic.
func PublishAmessage() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
	"type": "signup",
	"user": "john",
	"id": "1",
},
Topic: "user",

	})
	fmt.Println(rsp, err)
}
```
## Subscribe

Subscribe to messages for a given topic.


[https://m3o.com/event/api#Subscribe](https://m3o.com/event/api#Subscribe)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Subscribe to messages for a given topic.
func SubscribeToAtopic() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Subscribe(&event.SubscribeRequest{
		Topic: "user",

	})
	fmt.Println(rsp, err)
}
```
