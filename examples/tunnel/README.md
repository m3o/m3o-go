# Tunnel

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/tunnel/api](https://m3o.com/tunnel/api).

Endpoints:

## Send

Send a request through the tunnel


[https://m3o.com/tunnel/api#Send](https://m3o.com/tunnel/api#Send)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/tunnel"
)

// Send a request through the tunnel
func SendArequest() {
	tunnelService := tunnel.NewTunnelService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := tunnelService.Send(&tunnel.SendRequest{
		Url: "https://google.com",

	})
	fmt.Println(rsp, err)
	
}
```
