package price

import (
	"go.m3o.com/client"
)

type Price interface {
	Add(*AddRequest) (*AddResponse, error)
	Get(*GetRequest) (*GetResponse, error)
	Index(*IndexRequest) (*IndexResponse, error)
	List(*ListRequest) (*ListResponse, error)
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

type AddRequest struct {
	// currency e.g USD
	Currency string `json:"currency"`
	// name of the thing e.g bitcoin
	Name string `json:"name"`
	// price of the thing e.g 10001.00
	Price float64 `json:"price"`
	// symbol of value
	Symbol string `json:"symbol"`
}

type AddResponse struct {
	Value *Value `json:"value"`
}

type GetRequest struct {
	// currency to get
	Currency string `json:"currency"`
	// name of the value
	Name string `json:"name"`
	// symbol of value
	Symbol string `json:"symbol"`
}

type GetResponse struct {
	Values []Value `json:"values"`
}

type Index struct {
	// currency of value
	Currency string `json:"currency"`
	// name of item
	Name string `json:"name"`
	// symbol of item
	Symbol string `json:"symbol"`
}

type IndexRequest struct {
}

type IndexResponse struct {
	Index []Index `json:"index"`
}

type ListRequest struct {
	// currency to get
	Currency string `json:"currency"`
	// limit number of values
	Limit int32 `json:"limit"`
	// offset to read from
	Offset int32 `json:"offset"`
}

type ListResponse struct {
	Values []Value `json:"values"`
}

type Value struct {
	// currency of thing
	Currency string `json:"currency"`
	// name of thing
	Name string `json:"name"`
	// price of thing
	Price float64 `json:"price"`
	// symbol of value
	Symbol string `json:"symbol"`
	// time it was added
	Timestamp string `json:"timestamp"`
}
