# M3O Go Client

By default the client connects to api.m3o.com

```go
package main

import (
    "fmt"
    "os"

    "github.com/m3o/m3o-go/client"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Msg string `json:"msg"`
}

var (
	token, _ = os.Getenv("TOKEN")
)

func main() {
	c := client.NewClient(nil)

	// set your api token
	c.SetToken(token)

   	req := &Request{
		Name: "John",
	}
	var rsp Response

	if err := c.Call("helloworld", "call", req, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
```

## Streaming

The client supports streaming

```go
package main

import (
	"fmt"

	"github.com/m3o/m3o-go/client"
)

type Request struct {
	Count string `json:"count"`
}

type Response struct {
	Count string `json:"count"`
}

var (
	token, _ = os.Getenv("TOKEN")
)

func main() {
	c := client.NewClient(nil)

	// set your api token
	c.SetToken(token)
	
	stream, err := c.Stream("streams", "subscribe", Request{Count: "10"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var rsp Response
		if err := stream.Recv(&rsp); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("got", rsp.Count)
	}
}
```
