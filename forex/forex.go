package forex

import (
	"go.m3o.com/client"
)

type Forex interface {
	History(*HistoryRequest) (*HistoryResponse, error)
	Price(*PriceRequest) (*PriceResponse, error)
	Quote(*QuoteRequest) (*QuoteResponse, error)
}

func NewForexService(token string) *ForexService {
	return &ForexService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ForexService struct {
	client *client.Client
}

// Returns the data for the previous close
func (t *ForexService) History(request *HistoryRequest) (*HistoryResponse, error) {

	rsp := &HistoryResponse{}
	return rsp, t.client.Call("forex", "History", request, rsp)

}

// Get the latest price for a given forex ticker
func (t *ForexService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("forex", "Price", request, rsp)

}

// Get the latest quote for the forex
func (t *ForexService) Quote(request *QuoteRequest) (*QuoteResponse, error) {

	rsp := &QuoteResponse{}
	return rsp, t.client.Call("forex", "Quote", request, rsp)

}

type HistoryRequest struct {
	// the forex symbol e.g GBPUSD
	Symbol string `json:"symbol,omitempty"`
}

type HistoryResponse struct {
	// the close price
	Close float64 `json:"close,omitempty"`
	// the date
	Date string `json:"date,omitempty"`
	// the peak price
	High float64 `json:"high,omitempty"`
	// the low price
	Low float64 `json:"low,omitempty"`
	// the open price
	Open float64 `json:"open,omitempty"`
	// the forex symbol
	Symbol string `json:"symbol,omitempty"`
	// the volume
	Volume float64 `json:"volume,omitempty"`
}

type PriceRequest struct {
	// forex symbol e.g GBPUSD
	Symbol string `json:"symbol,omitempty"`
}

type PriceResponse struct {
	// the last price
	Price float64 `json:"price,omitempty"`
	// the forex symbol e.g GBPUSD
	Symbol string `json:"symbol,omitempty"`
}

type QuoteRequest struct {
	// the forex symbol e.g GBPUSD
	Symbol string `json:"symbol,omitempty"`
}

type QuoteResponse struct {
	// the asking price
	AskPrice float64 `json:"ask_price,omitempty"`
	// the bidding price
	BidPrice float64 `json:"bid_price,omitempty"`
	// the forex symbol
	Symbol string `json:"symbol,omitempty"`
	// the UTC timestamp of the quote
	Timestamp string `json:"timestamp,omitempty"`
}
