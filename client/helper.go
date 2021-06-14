package client

import (
	"bytes"
	"net/http"
	"net/url"
	"sort"
)

type CallObject struct {
	Priority int
	Service  string
	Endpoint string
	Request  map[string]interface{}
}

type CallObjectResponse struct {
	Priority int
	Service  string
	Endpoint string
	Request  map[string]interface{}
	Response interface{}
	Err      error
}

// RQ stores requests and RS stores requests and thier responses
type CallObjects struct {
	RS []CallObjectResponse
	RQ []CallObject
}

func NewCallObjects() *CallObjects {
	return &CallObjects{
		RS: []CallObjectResponse{},
		RQ: []CallObject{},
	}
}

// Implementing sort.Interface interface Len, Less and Swap
func (co *CallObjects) Len() int {
	return len(co.RQ)
}

func (co *CallObjects) Less(i, j int) bool {
	return co.RQ[i].Priority < co.RQ[j].Priority
}

func (co *CallObjects) Swap(i, j int) {
	co.RQ[i], co.RQ[j] = co.RQ[j], co.RQ[i]
}

// Put adds items(s)
func (co *CallObjects) Put(items ...CallObject) {
	co.RQ = append(co.RQ, items...)
}

// SortPriority sorts CallObjects by Priority
func (co *CallObjects) SortPriority() {
	sort.Sort(co)
}

// Get returns highest priority and remove it from CallObjects.RQ
func (co *CallObjects) Get() (bool, CallObject) {
	if len(co.RQ) == 0 {
		return false, CallObject{}
	}

	// last element
	i := len(co.RQ) - 1

	obj := co.RQ[i]

	co.RQ = co.RQ[0:i]

	return true, obj
}

// Empty checks if CallObjects.RQ is empty
func (co *CallObjects) Empty() bool {
	return len(co.RQ) == 0
}

// CreateRequest creates a request for a CallObject
func CreateRequest(co CallObject, client *Client) (*http.Request, error) {

	uri, err := url.Parse(client.options.Address)
	if err != nil {
		return nil, err
	}

	// set the url to go through the v1 api
	uri.Path = "/v1/" + co.Service + "/" + co.Endpoint

	b, err := marshalRequest(co.Service, co.Endpoint, co.Request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	// set the token if it exists
	if len(client.options.Token) > 0 {
		req.Header.Set("Authorization", "Bearer "+client.options.Token)
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
