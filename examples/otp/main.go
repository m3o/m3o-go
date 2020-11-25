package otp

import (
	"github.com/m3o/m3o-go/examples/otp/handler"
	"github.com/m3o/m3o-go/server"
	"github.com/m3o/m3o-go/sms"
	"github.com/m3o/m3o-go/store/keyvalue"
)

func main() {
	// register the handler
	server.RegisterHandler(handler.OTP{
		SMS:   sms.NewClient(),
		Store: keyvalue.NewClient(),
	})

	// run the server
	server.Run()
}
