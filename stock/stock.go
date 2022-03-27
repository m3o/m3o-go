package stock

import (
	"go.m3o.com/client"
)

type Stock interface {
	History(*HistoryRequest) (*HistoryResponse, error)
	OrderBook(*OrderBookRequest) (*OrderBookResponse, error)
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

// Get the historic order book and each trade by timestamp
func (t *StockService) OrderBook(request *OrderBookRequest) (*OrderBookResponse, error) {

	rsp := &OrderBookResponse{}
	return rsp, t.client.Call("stock", "OrderBook", request, rsp)

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

type Order struct {
	// the asking price
	AskPrice float64 `json:"ask_price,omitempty"`
	// the ask size
	AskSize int32 `json:"ask_size,omitempty"`
	// the bidding price
	BidPrice float64 `json:"bid_price,omitempty"`
	// the bid size
	BidSize int32 `json:"bid_size,omitempty"`
	// the UTC timestamp of the quote
	Timestamp string `json:"timestamp,omitempty"`
}

type OrderBookRequest struct {
	// the date in format YYYY-MM-dd
	Date string `json:"date,omitempty"`
	// optional RFC3339Nano end time e.g 2006-01-02T15:04:05.999999999Z07:00
	End string `json:"end,omitempty"`
	// limit number of prices
	Limit int32 `json:"limit,omitempty"`
	// optional RFC3339Nano start time e.g 2006-01-02T15:04:05.999999999Z07:00
	Start string `json:"start,omitempty"`
	// stock to retrieve e.g AAPL
	Stock string `json:"stock,omitempty"`
}

type OrderBookResponse struct {
	// date of the request
	Date string `json:"date,omitempty"`
	// list of orders
	Orders []Order `json:"orders,omitempty"`
	// the stock symbol
	Symbol string `json:"symbol,omitempty"`
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
