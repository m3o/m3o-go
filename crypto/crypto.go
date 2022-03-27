package crypto

import (
	"go.m3o.com/client"
)

type Crypto interface {
	History(*HistoryRequest) (*HistoryResponse, error)
	News(*NewsRequest) (*NewsResponse, error)
	Price(*PriceRequest) (*PriceResponse, error)
	Quote(*QuoteRequest) (*QuoteResponse, error)
	Symbols(*SymbolsRequest) (*SymbolsResponse, error)
}

func NewCryptoService(token string) *CryptoService {
	return &CryptoService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CryptoService struct {
	client *client.Client
}

// Returns the history for the previous close
func (t *CryptoService) History(request *HistoryRequest) (*HistoryResponse, error) {

	rsp := &HistoryResponse{}
	return rsp, t.client.Call("crypto", "History", request, rsp)

}

// Get news related to a currency
func (t *CryptoService) News(request *NewsRequest) (*NewsResponse, error) {

	rsp := &NewsResponse{}
	return rsp, t.client.Call("crypto", "News", request, rsp)

}

// Get the last price for a given crypto ticker
func (t *CryptoService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("crypto", "Price", request, rsp)

}

// Get the last quote for a given crypto ticker
func (t *CryptoService) Quote(request *QuoteRequest) (*QuoteResponse, error) {

	rsp := &QuoteResponse{}
	return rsp, t.client.Call("crypto", "Quote", request, rsp)

}

// Returns the full list of supported symbols
func (t *CryptoService) Symbols(request *SymbolsRequest) (*SymbolsResponse, error) {

	rsp := &SymbolsResponse{}
	return rsp, t.client.Call("crypto", "Symbols", request, rsp)

}

type Article struct {
	// the date published
	Date string `json:"date,omitempty"`
	// its description
	Description string `json:"description,omitempty"`
	// the source
	Source string `json:"source,omitempty"`
	// title of the article
	Title string `json:"title,omitempty"`
	// the source url
	Url string `json:"url,omitempty"`
}

type HistoryRequest struct {
	// the crypto symbol e.g BTCUSD
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
	// the crypto symbol
	Symbol string `json:"symbol,omitempty"`
	// the volume
	Volume float64 `json:"volume,omitempty"`
}

type NewsRequest struct {
	// cryptocurrency ticker to request news for e.g BTC
	Symbol string `json:"symbol,omitempty"`
}

type NewsResponse struct {
	// list of articles
	Articles []Article `json:"articles,omitempty"`
	// symbol requested for
	Symbol string `json:"symbol,omitempty"`
}

type PriceRequest struct {
	// the crypto symbol e.g BTCUSD
	Symbol string `json:"symbol,omitempty"`
}

type PriceResponse struct {
	// the last price
	Price float64 `json:"price,omitempty"`
	// the crypto symbol e.g BTCUSD
	Symbol string `json:"symbol,omitempty"`
}

type QuoteRequest struct {
	// the crypto symbol e.g BTCUSD
	Symbol string `json:"symbol,omitempty"`
}

type QuoteResponse struct {
	// the asking price
	AskPrice float64 `json:"ask_price,omitempty"`
	// the ask size
	AskSize float64 `json:"ask_size,omitempty"`
	// the bidding price
	BidPrice float64 `json:"bid_price,omitempty"`
	// the bid size
	BidSize float64 `json:"bid_size,omitempty"`
	// the crypto symbol
	Symbol string `json:"symbol,omitempty"`
	// the UTC timestamp of the quote
	Timestamp string `json:"timestamp,omitempty"`
}

type Symbol struct {
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type SymbolsRequest struct {
}

type SymbolsResponse struct {
	Symbols []Symbol `json:"symbols,omitempty"`
}
