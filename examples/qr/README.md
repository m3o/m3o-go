# Qr

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/qr/api](https://m3o.com/qr/api).

Endpoints:

## Codes

List your QR codes


[https://m3o.com/qr/api#Codes](https://m3o.com/qr/api#Codes)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/qr"
)

// List your QR codes
func ListCodes() {
	qrService := qr.NewQrService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := qrService.Codes(&qr.CodesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate a QR code with a specific text and size


[https://m3o.com/qr/api#Generate](https://m3o.com/qr/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/qr"
)

// Generate a QR code with a specific text and size
func GenerateAqrCode() {
	qrService := qr.NewQrService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := qrService.Generate(&qr.GenerateRequest{
		Size: 300,
Text: "https://m3o.com/qr",

	})
	fmt.Println(rsp, err)
	
}
```
