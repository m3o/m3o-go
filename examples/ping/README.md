# Ping

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/ping/api](https://m3o.com/ping/api).

Endpoints:

## Url

Ping a HTTP URL


[https://m3o.com/ping/api#Url](https://m3o.com/ping/api#Url)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping a HTTP URL
func CheckAurl() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Url(&ping.UrlRequest{
		Address: "google.com",

	})
	fmt.Println(rsp, err)
	
}
```
## Ip

Ping an IP address


[https://m3o.com/ping/api#Ip](https://m3o.com/ping/api#Ip)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping an IP address
func PingAnIp() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Ip(&ping.IpRequest{
		Address: "google.com",

	})
	fmt.Println(rsp, err)
	
}
```
## Tcp

Ping a TCP port to check if it's open


[https://m3o.com/ping/api#Tcp](https://m3o.com/ping/api#Tcp)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping a TCP port to check if it's open
func DialAtcpAddress() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Tcp(&ping.TcpRequest{
		Address: "google.com:80",

	})
	fmt.Println(rsp, err)
	
}
```
