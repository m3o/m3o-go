# Mq

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/mq/api](https://m3o.com/mq/api).

Endpoints:

## Subscribe

Subscribe to messages for a given topic.


[https://m3o.com/mq/api#Subscribe](https://m3o.com/mq/api#Subscribe)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Subscribe to messages for a given topic.
func SubscribeToAtopic() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := mqService.Subscribe(&mq.SubscribeRequest{
		Topic: "events",

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
## Publish

Publish a message. Specify a topic to group messages for a specific topic.


[https://m3o.com/mq/api#Publish](https://m3o.com/mq/api#Publish)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/mq"
)

// Publish a message. Specify a topic to group messages for a specific topic.
func PublishAmessage() {
	mqService := mq.NewMqService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := mqService.Publish(&mq.PublishRequest{
		Message: map[string]interface{}{
	"user": "john",
	"id": "1",
	"type": "signup",
},
Topic: "events",

	})
	fmt.Println(rsp, err)
	
}
```
