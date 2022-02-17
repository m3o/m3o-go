# Helloworld

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/helloworld/api](https://m3o.com/helloworld/api).

Endpoints:

## Stream

Stream returns a stream of "Hello $name" responses


[https://m3o.com/helloworld/api#Stream](https://m3o.com/helloworld/api#Stream)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Stream returns a stream of "Hello $name" responses
func StreamsResponsesFromTheServerUsingWebsockets() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := helloworldService.Stream(&helloworld.StreamRequest{
		Messages: 10,
Name: "John",

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
## Call

Call returns a personalised "Hello $name" response


[https://m3o.com/helloworld/api#Call](https://m3o.com/helloworld/api#Call)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Call returns a personalised "Hello $name" response
func CallTheHelloworldService() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := helloworldService.Call(&helloworld.CallRequest{
		Name: "John",

	})
	fmt.Println(rsp, err)
	
}
```
