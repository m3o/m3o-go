# Address

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Address/api](https://m3o.com/Address/api).

Endpoints:

## LookupPostcode

Lookup a list of UK addresses by postcode


[https://m3o.com/address/api#LookupPostcode](https://m3o.com/address/api#LookupPostcode)

```go
package example

import(
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
```
