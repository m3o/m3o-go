package handler

import (
	"crypto/rand"

	otp "github.com/m3o/m3o-go/examples/otp/proto"
	"github.com/m3o/m3o-go/context"
	"github.com/m3o/m3o-go/errors"
	"github.com/m3o/m3o-go/sms"
	"github.com/m3o/m3o-go/store/keyvalue"
)

const (
	otpChars  = "1234567890"
	otpLength = 6
	otpTTL = time.Minute * 5
)

var (
	errMissingPhoneNumber = errors.BadRequest("Missing PhoneNumber")
	errMissingCode 				= errors.BadRequest("Missing Code")
	errInvalidCode 				= errors.BadRequest("Invalid Code")
)

// OTP satisfies the OTP handler interface. Note: this handler is an exampple and not intended for 
// production use - the handler does not enforce verification limits so it would be possible to brute
// force the verification process. It is advised to use a retry limit of 3 in real-world applications.
type OTP struct {
	SMS   sms.Service
	Store keyvalue.Store
}

// Send a one time password to the phone number provided
func (o *OTP) Send(ctx context.Context, req *otp.SendRequest) (*otp.SendResponse, error) {
	// validate the request
	if len(req.PhoneNumber) == 0 {
		return nil, errMissingPhoneNumber
	}

	// generate and send the OPT
	otp := generateOTP()
	msg := fmt.Sprintf("Your M3O verification code is %v", otp
	if err := o.SMS.Send(req.PhoneNumber, msg); err != nil {
		return nil, err
	}

	// write the code to the store
	err := o.Store.Write(req.PhoneNumber, otp, keyvalue.WriteOptions{
		Expiry: time.Now().Add(otpTTL),
	})
	if err != nil {
		return nil, errors.InternalServerError("Error writing OTP to the store")
	}

	// return no error, indicating the verification code send successfully
	return &otp.SendResponse{}, nil
}

// Validate a one time password
func (o *OTP) Verify(ctx context.Context, req *otp.VerifyRequest) (*otp.VerifyResponse, error) {
	// validate the request
	if len(req.PhoneNumber) == 0 {
		return nil, errMissingPhoneNumber
	}
	if len(req.Code) == 0 {
		return nil, errMissingCode
	}

	// lookup the phone number from the store
	rec, err := o.Store.Read(req.PhoneNumber)
	if err == keyvalue.ErrRecordNotFound {
		return nil, errMissingCode
	} else if err != nil {
		return nil, errors.InternalServerError("Error reading OTP from the store")
	}

	// ensure the codes match
	if rec.Value != req.Code {
		return nil, errInvalidCode
	}

	// return no error, indicating the verification was successful
	return &otp.VerifyResponse{}, nil
}

func generateOTP() (string, error) {
	buffer := make([]byte, otpLength)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < otpLength; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
