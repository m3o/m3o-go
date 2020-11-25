package handler

import (
	"testing"

	"github.com/m3o/m3o-go/sms"
	"github.com/m3o/m3o-go/store/keyvalue"
)

func TestHandler(t *testing.T) {
	h := OTP{
		SMS:   sms.NewMock(),
		Store: keyvalue.NewMock(),
	}

	t.Run("Send", func(t *testing.T) {

	})

	t.Run("Verify", func(t *testing.T) {

	})
}
