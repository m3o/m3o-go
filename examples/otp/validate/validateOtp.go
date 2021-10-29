package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/otp"
)

// Validate the OTP code
func ValidateOtp() {
	otpService := otp.NewOtpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := otpService.Validate(&otp.ValidateRequest{
		Code: "656211",
		Id:   "asim@example.com",
	})
	fmt.Println(rsp, err)
}
