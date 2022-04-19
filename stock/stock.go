package stock

import (
	"go.m3o.com/client"
)

type Stock interface {
	History(*HistoryRequest) (*HistoryResponse, error)
	Price(*PriceRequest) (*PriceResponse, error)
	Quote(*QuoteRequest) (*QuoteResponse, error)
}

func NewStockService(token string) *StockService {
	return &StockService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type StockService struct {
	client *client.Client
}

// Get the historic open-close for a given day
func (t *StockService) History(request *HistoryRequest) (*HistoryResponse, error) {

	rsp := &HistoryResponse{}
	return rsp, t.client.Call("stock", "History", request, rsp)

}

// Get the last price for a given stock ticker
func (t *StockService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("stock", "Price", request, rsp)

}

// Get the last quote for the stock
func (t *StockService) Quote(request *QuoteRequest) (*QuoteResponse, error) {

	rsp := &QuoteResponse{}
	return rsp, t.client.Call("stock", "Quote", request, rsp)

}

type HistoryRequest struct {
	// date to retrieve as YYYY-MM-DD
	Date string `json:"date,omitempty"`
	// the stock symbol e.g AAPL
	Stock string `json:"stock,omitempty"`
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
	// the stock symbol
	Symbol string `json:"symbol,omitempty"`
	// the volume
	Volume int32 `json:"volume,omitempty"`
}

type PriceRequest struct {
	// stock symbol e.g AAPL
	Symbol string `json:"symbol,omitempty"`
}

type PriceResponse struct {
	// the last price
	Price float64 `json:"price,omitempty"`
	// the stock symbol e.g AAPL
	Symbol string `json:"symbol,omitempty"`
}

type QuoteRequest struct {
	// the stock symbol e.g AAPL
	Symbol string `json:"symbol,omitempty"`
}

type QuoteResponse struct {
	// the asking price
	AskPrice float64 `json:"ask_price,omitempty"`
	// the ask size
	AskSize int32 `json:"ask_size,omitempty"`
	// the bidding price
	BidPrice float64 `json:"bid_price,omitempty"`
	// the bid size
	BidSize int32 `json:"bid_size,omitempty"`
	// the stock symbol
	Symbol string `json:"symbol,omitempty"`
	// the UTC timestamp of the quote
	Timestamp string `json:"timestamp,omitempty"`
}
