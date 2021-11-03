package main

import (
	"fmt"
	"os"

	"go.m3o.com/otp"
)

// Generate an OTP (one time pass) code
func main() {
	otpService := otp.NewOtpService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := otpService.Generate(&otp.GenerateRequest{
		Id: "asim@example.com",
	})
	fmt.Println(rsp, err)

}
