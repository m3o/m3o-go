package bitcoin

import (
	"go.m3o.com/client"
)

type Bitcoin interface {
	Balance(*BalanceRequest) (*BalanceResponse, error)
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

// Get the BTC balance of an address
func (t *BitcoinService) Balance(request *BalanceRequest) (*BalanceResponse, error) {

	rsp := &BalanceResponse{}
	return rsp, t.client.Call("bitcoin", "Balance", request, rsp)

}

// Get the price of bitcoin
func (t *BitcoinService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("bitcoin", "Price", request, rsp)

}

type BalanceRequest struct {
	// address to lookup
	Address string `json:"address,omitempty"`
}

type BalanceResponse struct {
	// total BTC as satoshis
	Balance int64 `json:"balance,string,omitempty"`
}

type PriceRequest struct {
	// symbol to use e.g BTCUSD
	Symbol string `json:"symbol,omitempty"`
}

type PriceResponse struct {
	// The price of bitcoin
	Price float64 `json:"price,omitempty"`
	// The symbol of pricing e.g BTCUSD
	Symbol string `json:"symbol,omitempty"`
}
