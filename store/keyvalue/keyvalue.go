// Package keyvalue provides a key value store
package keyvalue

import (
	"errors"
	"time"
)

var (
	// ErrRecordNotFound is returned when trying to read a record which does not exist
	ErrRecordNotFound = errors.New("Record not found")
)

// NewClient returns an RPC client for the KeyValue store. It will communicate with the M3O KeyValue
// service.
func NewClient() Store {

}

// NewMock returns an in-memory mock for the KeyValue store. It is designed for use in tests.
func NewMock() Store {

}

// Store is an interface providing key value storage
type Store interface {
	Read(key string) (Record, error)
	List(ListOptions) ([]Record, error)
	Write(key, value string, opts WriteOptions) error
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
}

// Record returned from the key value store
type Record struct {
	Key    string
	Value  string
	Expiry time.Time
}
