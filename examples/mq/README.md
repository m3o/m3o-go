# Mq

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Mq/api](https://m3o.com/Mq/api).

Endpoints:

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
	"id": "1",
	"type": "signup",
	"user": "john",
},
Topic: "events",

	})
	fmt.Println(rsp, err)
}
```
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
	rsp, err := mqService.Subscribe(&mq.SubscribeRequest{
		Topic: "events",

	})
	fmt.Println(rsp, err)
}
```
