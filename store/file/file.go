// Package file provides a file store
package file

import (
	"io"
	"time"
)

// NewClient returns an RPC client for the file store. It will communicate with the M3O file
// service.
func NewClient() File {

}

// NewMock returns an in-memory mock for the file store. It is designed for use in tests.
func NewMock() File {

}

// File is an interface providing file storage
type File interface {
	Read(key string) (Result, error)
	List(ListOptions) ([]Result, error)
	Write(key string, value io.Reader, opts WriteOptions) error
	Delete(key string) error
}

// ListOptions are used to filter results of the list operation
type ListOptions struct {
	// Prefix limits the results to those where the key has the given prefix
	Prefix string
}

// WriteOptions are provided when writing a file to the store
type WriteOptions struct {
	// Expiry sets the time at which the file should be deleted
	Expiry time.Time
	// Encoding of the file, e.g. PDF
	Encoding string
}

// Result returned from the key value store
type Result struct {
	Key      string
	Value    io.Reader
	Expiry   time.Time
	Encoding string
}
