package example

import (
	"fmt"
	"os"

	"go.m3o.com/ip"
)

// Lookup the geolocation information for an IP address
func LookupIpInfo() {
	ipService := ip.NewIpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := ipService.Lookup(&ip.LookupRequest{
		Ip: "93.148.214.31",
	})
	fmt.Println(rsp, err)
}
