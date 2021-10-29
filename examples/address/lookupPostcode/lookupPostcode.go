package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/address"
)

// Lookup a list of UK addresses by postcode
func LookupPostcode() {
	addressService := address.NewAddressService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := addressService.LookupPostcode(&address.LookupPostcodeRequest{
		Postcode: "SW1A 2AA",
	})
	fmt.Println(rsp, err)
}
