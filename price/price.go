package price

import (
	"go.m3o.com/client"
)

type Price interface {
	Add(*AddRequest) (*AddResponse, error)
	Get(*GetRequest) (*GetResponse, error)
	Index(*IndexRequest) (*IndexResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Report(*ReportRequest) (*ReportResponse, error)
}

func NewPriceService(token string) *PriceService {
	return &PriceService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PriceService struct {
	client *client.Client
}

// Add a price
func (t *PriceService) Add(request *AddRequest) (*AddResponse, error) {

	rsp := &AddResponse{}
	return rsp, t.client.Call("price", "Add", request, rsp)

}

// Get the price of anything
func (t *PriceService) Get(request *GetRequest) (*GetResponse, error) {

	rsp := &GetResponse{}
	return rsp, t.client.Call("price", "Get", request, rsp)

}

// Get the index for available prices
func (t *PriceService) Index(request *IndexRequest) (*IndexResponse, error) {

	rsp := &IndexResponse{}
	return rsp, t.client.Call("price", "Index", request, rsp)

}

// List prices for a given currency
func (t *PriceService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("price", "List", request, rsp)

}

// Report an invalid price
func (t *PriceService) Report(request *ReportRequest) (*ReportResponse, error) {

	rsp := &ReportResponse{}
	return rsp, t.client.Call("price", "Report", request, rsp)

}

type AddRequest struct {
	// author of the price
	Author string `json:"author,omitempty"`
	// currency e.g USD
	Currency string `json:"currency,omitempty"`
	// name of the thing e.g bitcoin
	Name string `json:"name,omitempty"`
	// price of the thing e.g 10001.00
	Price float64 `json:"price,omitempty"`
	// source of the price
	Source string `json:"source,omitempty"`
	// symbol of value
	Symbol string `json:"symbol,omitempty"`
}

type AddResponse struct {
	Value *Value `json:"value,omitempty"`
}

type GetRequest struct {
	// currency to get
	Currency string `json:"currency,omitempty"`
	// name of the value
	Name string `json:"name,omitempty"`
	// symbol of value
	Symbol string `json:"symbol,omitempty"`
}

type GetResponse struct {
	Values []Value `json:"values,omitempty"`
}

type Index struct {
	// currency of value
	Currency string `json:"currency,omitempty"`
	// name of item
	Name string `json:"name,omitempty"`
	// symbol of item
	Symbol string `json:"symbol,omitempty"`
}

type IndexRequest struct {
}

type IndexResponse struct {
	Index []Index `json:"index,omitempty"`
}

type ListRequest struct {
	// currency to get
	Currency string `json:"currency,omitempty"`
	// limit number of values
	Limit int32 `json:"limit,omitempty"`
	// offset to read from
	Offset int32 `json:"offset,omitempty"`
}

type ListResponse struct {
	Values []Value `json:"values,omitempty"`
}

type Report struct {
	Author  string `json:"author,omitempty"`
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
	Symbol  string `json:"symbol,omitempty"`
}

type ReportRequest struct {
	// additional comment
	Comment string `json:"comment,omitempty"`
	// name of value
	Name string `json:"name,omitempty"`
	// symbol of value
	Symbol string `json:"symbol,omitempty"`
}

type ReportResponse struct {
}

type Value struct {
	// who added it
	Author string `json:"author,omitempty"`
	// currency of thing
	Currency string `json:"currency,omitempty"`
	// name of thing
	Name string `json:"name,omitempty"`
	// price of thing
	Price float64 `json:"price,omitempty"`
	// where it came from
	Source string `json:"source,omitempty"`
	// symbol of value
	Symbol string `json:"symbol,omitempty"`
	// time it was added
	Timestamp string `json:"timestamp,omitempty"`
}
