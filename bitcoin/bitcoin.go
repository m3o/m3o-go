package bitcoin

import (
	"go.m3o.com/client"
)

type Bitcoin interface {
	Price(*PriceRequest) (*PriceResponse, error)
}

func NewBitcoinService(token string) *BitcoinService {
	return &BitcoinService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type BitcoinService struct {
	client *client.Client
}

// Get the price of bitcoin
func (t *BitcoinService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("bitcoin", "Price", request, rsp)

}

type PriceRequest struct {
	// symbol to use e.g BTCUSD
	Symbol string `json:"symbol"`
}

type PriceResponse struct {
	// The price of bitcoin
	Price float64 `json:"price"`
	// The symbol of pricing e.g BTCUSD
	Symbol string `json:"symbol"`
}
