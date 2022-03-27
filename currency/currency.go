package currency

import (
	"go.m3o.com/client"
)

type Currency interface {
	Codes(*CodesRequest) (*CodesResponse, error)
	Convert(*ConvertRequest) (*ConvertResponse, error)
	History(*HistoryRequest) (*HistoryResponse, error)
	Rates(*RatesRequest) (*RatesResponse, error)
}

func NewCurrencyService(token string) *CurrencyService {
	return &CurrencyService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CurrencyService struct {
	client *client.Client
}

// Codes returns the supported currency codes for the API
func (t *CurrencyService) Codes(request *CodesRequest) (*CodesResponse, error) {

	rsp := &CodesResponse{}
	return rsp, t.client.Call("currency", "Codes", request, rsp)

}

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
func (t *CurrencyService) Convert(request *ConvertRequest) (*ConvertResponse, error) {

	rsp := &ConvertResponse{}
	return rsp, t.client.Call("currency", "Convert", request, rsp)

}

// Returns the historic rates for a currency on a given date
func (t *CurrencyService) History(request *HistoryRequest) (*HistoryResponse, error) {

	rsp := &HistoryResponse{}
	return rsp, t.client.Call("currency", "History", request, rsp)

}

// Rates returns the currency rates for a given code e.g USD
func (t *CurrencyService) Rates(request *RatesRequest) (*RatesResponse, error) {

	rsp := &RatesResponse{}
	return rsp, t.client.Call("currency", "Rates", request, rsp)

}

type Code struct {
	// e.g United States Dollar
	Currency string `json:"currency,omitempty"`
	// e.g USD
	Name string `json:"name,omitempty"`
}

type CodesRequest struct {
}

type CodesResponse struct {
	Codes []Code `json:"codes,omitempty"`
}

type ConvertRequest struct {
	// optional amount to convert e.g 10.0
	Amount float64 `json:"amount,omitempty"`
	// base code to convert from e.g USD
	From string `json:"from,omitempty"`
	// target code to convert to e.g GBP
	To string `json:"to,omitempty"`
}

type ConvertResponse struct {
	// converted amount e.g 7.10
	Amount float64 `json:"amount,omitempty"`
	// the base code e.g USD
	From string `json:"from,omitempty"`
	// conversion rate e.g 0.71
	Rate float64 `json:"rate,omitempty"`
	// the target code e.g GBP
	To string `json:"to,omitempty"`
}

type HistoryRequest struct {
	// currency code e.g USD
	Code string `json:"code,omitempty"`
	// date formatted as YYYY-MM-DD
	Date string `json:"date,omitempty"`
}

type HistoryResponse struct {
	// The code of the request
	Code string `json:"code,omitempty"`
	// The date requested
	Date string `json:"date,omitempty"`
	// The rate for the day as code:rate
	Rates map[string]float64 `json:"rates,omitempty"`
}

type RatesRequest struct {
	// The currency code to get rates for e.g USD
	Code string `json:"code,omitempty"`
}

type RatesResponse struct {
	// The code requested e.g USD
	Code string `json:"code,omitempty"`
	// The rates for the given code as key-value pairs code:rate
	Rates map[string]float64 `json:"rates,omitempty"`
}
