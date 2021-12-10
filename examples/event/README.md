# Event

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Event/api](https://m3o.com/Event/api).

Endpoints:

## Publish

Publish a event to the event stream.


[https://m3o.com/event/api#Publish](https://m3o.com/event/api#Publish)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Publish a event to the event stream.
func PublishAnEvent() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
	"id": "1",
	"type": "signup",
	"user": "john",
},
Topic: "user",

	})
	fmt.Println(rsp, err)
	
}
```
## Consume

Consume events from a given topic.


[https://m3o.com/event/api#Consume](https://m3o.com/event/api#Consume)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Consume events from a given topic.
func ConsumeFromAtopic() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := eventService.Consume(&event.ConsumeRequest{
		Topic: "user",

	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
			rsp, err := stream.Recv()
			if err != nil {
					fmt.Println(err)
					return
			}

			fmt.Println(rsp)
	}
}
```
## Read

Read stored events


[https://m3o.com/event/api#Read](https://m3o.com/event/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Read stored events
func ReadEventsOnAtopic() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Read(&event.ReadRequest{
		Topic: "user",

	})
	fmt.Println(rsp, err)
	
}
```
