package bitcoin

import (
	"go.m3o.com/client"
)

type Bitcoin interface {
	Balance(*BalanceRequest) (*BalanceResponse, error)
	Lookup(*LookupRequest) (*LookupResponse, error)
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

// Get details for a bitcoin address
func (t *BitcoinService) Lookup(request *LookupRequest) (*LookupResponse, error) {

	rsp := &LookupResponse{}
	return rsp, t.client.Call("bitcoin", "Lookup", request, rsp)

}

// Get the price of bitcoin
func (t *BitcoinService) Price(request *PriceRequest) (*PriceResponse, error) {

	rsp := &PriceResponse{}
	return rsp, t.client.Call("bitcoin", "Price", request, rsp)

}

// Get transaction details by hash
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

type LookupRequest struct {
	// bitcoin address
	Address string `json:"address,omitempty"`
	// limit num transactions (max: 50)
	Limit int32 `json:"limit,omitempty"`
	// offset transactions
	Offset int32 `json:"offset,omitempty"`
}

type LookupResponse struct {
	// address requested
	Address string `json:"address,omitempty"`
	// final balanace
	FinalBalance int64 `json:"final_balance,string,omitempty"`
	// hash160
	Hash string `json:"hash,omitempty"`
	// number of transactions
	NumTx int64 `json:"num_tx,string,omitempty"`
	// number of unredeemed
	NumUnredeemed int64 `json:"num_unredeemed,string,omitempty"`
	// total received
	TotalReceived int64 `json:"total_received,string,omitempty"`
	// total sent
	TotalSent int64 `json:"total_sent,string,omitempty"`
	// list of transactions
	Transactions []Transaction `json:"transactions,omitempty"`
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
	N       int64  `json:"n,string,omitempty"`
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

type Transaction struct {
	// balance after transaction
	Balance int64 `json:"balance,string,omitempty"`
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
	// result of transaction
	Result int64 `json:"result,string,omitempty"`
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
