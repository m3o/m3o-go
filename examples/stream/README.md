# Stream

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Stream/api](https://m3o.com/Stream/api).

Endpoints:

## SendMessage

Send a message to the stream.


[https://m3o.com/stream/api#SendMessage](https://m3o.com/stream/api#SendMessage)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// Send a message to the stream.
func SendMessage() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.SendMessage(&stream.SendMessageRequest{
		Channel: "general",
Text: "Hey checkout this tweet https://twitter.com/m3oservices/status/1455291054295498752",

	})
	fmt.Println(rsp, err)
	
}
```
## ListMessages

List messages for a given channel


[https://m3o.com/stream/api#ListMessages](https://m3o.com/stream/api#ListMessages)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// List messages for a given channel
func ListMessages() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.ListMessages(&stream.ListMessagesRequest{
		Channel: "general",

	})
	fmt.Println(rsp, err)
	
}
```
## ListChannels

List all the active channels


[https://m3o.com/stream/api#ListChannels](https://m3o.com/stream/api#ListChannels)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// List all the active channels
func ListChannels() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.ListChannels(&stream.ListChannelsRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## CreateChannel

Create a channel with a given name and description. Channels are created automatically but
this allows you to specify a description that's persisted for the lifetime of the channel.


[https://m3o.com/stream/api#CreateChannel](https://m3o.com/stream/api#CreateChannel)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// Create a channel with a given name and description. Channels are created automatically but
// this allows you to specify a description that's persisted for the lifetime of the channel.
func CreateChannel() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.CreateChannel(&stream.CreateChannelRequest{
		Description: "The channel for all things",
Name: "general",

	})
	fmt.Println(rsp, err)
	
}
```
