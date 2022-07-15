package bitcoin

import (
	"go.m3o.com/client"
)

type Bitcoin interface {
	Balance(*BalanceRequest) (*BalanceResponse, error)
	Price(*PriceRequest) (*PriceResponse, error)
	Transaction(*TransactionRequest) (*TransactionResponse, error)
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

// Get the details of a transaction
func (t *BitcoinService) Transaction(request *TransactionRequest) (*TransactionResponse, error) {

	rsp := &TransactionResponse{}
	return rsp, t.client.Call("bitcoin", "Transaction", request, rsp)

}

type BalanceRequest struct {
	// address to lookup
	Address string `json:"address,omitempty"`
}

type BalanceResponse struct {
	// total BTC as satoshis
	Balance int64 `json:"balance,string,omitempty"`
}

type Input struct {
	PrevOut *Prev  `json:"prev_out,omitempty"`
	Script  string `json:"script,omitempty"`
}

type Output struct {
	Address string `json:"address,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Script  string `json:"script,omitempty"`
	Spent   bool   `json:"spent,omitempty"`
	TxIndex int64  `json:"tx_index,string,omitempty"`
	Value   int64  `json:"value,string,omitempty"`
}

type Prev struct {
	Address string `json:"address,omitempty"`
	Hash    string `json:"hash,omitempty"`
	N       string `json:"n,omitempty"`
	Script  string `json:"script,omitempty"`
	Spent   bool   `json:"spent,omitempty"`
	TxIndex int64  `json:"tx_index,string,omitempty"`
	Value   int64  `json:"value,string,omitempty"`
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

type TransactionRequest struct {
	// the transaction hash
	Hash string `json:"hash,omitempty"`
}

type TransactionResponse struct {
	// block height
	BlockHeight int64 `json:"block_height,string,omitempty"`
	// blck index
	BlockIndex int64 `json:"block_index,string,omitempty"`
	// double spend
	DoubleSpend bool `json:"double_spend,omitempty"`
	// fees
	Fee int64 `json:"fee,string,omitempty"`
	// transaction hash
	Hash string `json:"hash,omitempty"`
	// inputs
	Inputs []Input `json:"inputs,omitempty"`
	// lock time
	LockTime int64 `json:"lock_time,string,omitempty"`
	// outputs
	Outputs []Output `json:"outputs,omitempty"`
	// relay
	Relay string `json:"relay,omitempty"`
	// transaction size
	Size int64 `json:"size,string,omitempty"`
	// tx index
	TxIndex int64 `json:"tx_index,string,omitempty"`
	// the version
	Version int64 `json:"version,string,omitempty"`
	// vin
	VinSz int64 `json:"vin_sz,string,omitempty"`
	// vout
	VoutSz int64 `json:"vout_sz,string,omitempty"`
	// weight
	Weight int64 `json:"weight,string,omitempty"`
}
