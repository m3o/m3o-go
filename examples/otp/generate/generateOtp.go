package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/otp"
)

// Generate an OTP (one time pass) code
func GenerateOtp() {
	otpService := otp.NewOtpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := otpService.Generate(&otp.GenerateRequest{
		Id: "asim@example.com",
	})
	fmt.Println(rsp, err)
}
