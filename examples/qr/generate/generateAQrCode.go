package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/qr"
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
