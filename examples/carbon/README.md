# Carbon

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/carbon/api](https://m3o.com/carbon/api).

Endpoints:

## Offset

Purchase 1KG (0.001 tonne) of carbon offsets in a single request


[https://m3o.com/carbon/api#Offset](https://m3o.com/carbon/api#Offset)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/carbon"
)

// Purchase 1KG (0.001 tonne) of carbon offsets in a single request
func OffsetCarbon() {
	carbonService := carbon.NewCarbonService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := carbonService.Offset(&carbon.OffsetRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
