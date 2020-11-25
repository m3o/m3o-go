// Package sms provides the ability to send a text message
package sms

import "github.com/m3o/m3o-go/errors"

// NewClient returns an RPC client for the SMS service. It will communicate with the M3O SMS service
func NewClient() Service {

}

// NewMock returns an mock SMS service designed for usage with test
func NewMock() Service {

}

var (
	// ErrMissingNumber is returned when a blank number is passed to Send
	ErrMissingNumber = errors.BadRequest("Missing Number")
	// ErrInvalidNumber is returned when the number passed to Send could not be parsed
	ErrInvalidNumber = errors.BadRequest("Invalid Number")
	// ErrMissingMessage is returned when a blank message is passed to Send
	ErrMissingMessage = errors.BadRequest("Missing Message")
)

// Service is an interface providing SMS
type Service interface {
	Send(number, message string) error
}
