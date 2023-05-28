# Dns

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/dns/api](https://m3o.com/dns/api).

Endpoints:

## Whois

Check who owns a domain


[https://m3o.com/dns/api#Whois](https://m3o.com/dns/api#Whois)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/dns"
)

// Check who owns a domain
func WhoisQuery() {
	dnsService := dns.NewDnsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dnsService.Whois(&dns.WhoisRequest{
		Domain: "x.com",

	})
	fmt.Println(rsp, err)
	
}
```
## Query

Query an addresss


[https://m3o.com/dns/api#Query](https://m3o.com/dns/api#Query)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/dns"
)

// Query an addresss
func MakeAdnsQuery() {
	dnsService := dns.NewDnsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dnsService.Query(&dns.QueryRequest{
		Name: "google.com",

	})
	fmt.Println(rsp, err)
	
}
```
