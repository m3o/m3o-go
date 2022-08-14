package ethereum

import (
	"go.m3o.com/client"
)

type Ethereum interface {
	Balance(*BalanceRequest) (*BalanceResponse, error)
	Broadcast(*BroadcastRequest) (*BroadcastResponse, error)
	Transaction(*TransactionRequest) (*TransactionResponse, error)
}

func NewEthereumService(token string) *EthereumService {
	return &EthereumService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EthereumService struct {
	client *client.Client
}

// Get the balance of an ethereum wallet
func (t *EthereumService) Balance(request *BalanceRequest) (*BalanceResponse, error) {

	rsp := &BalanceResponse{}
	return rsp, t.client.Call("ethereum", "Balance", request, rsp)

}

// Broadcast presigned transaction to ethereum network
func (t *EthereumService) Broadcast(request *BroadcastRequest) (*BroadcastResponse, error) {

	rsp := &BroadcastResponse{}
	return rsp, t.client.Call("ethereum", "Broadcast", request, rsp)

}

// Get transaction details by hash
func (t *EthereumService) Transaction(request *TransactionRequest) (*TransactionResponse, error) {

	rsp := &TransactionResponse{}
	return rsp, t.client.Call("ethereum", "Transaction", request, rsp)

}

type BalanceRequest struct {
	// address of wallet
	Address string `json:"address,omitempty"`
}

type BalanceResponse struct {
	// the account balance (in wei)
	Balance int64 `json:"balance,string,omitempty"`
}

type BroadcastRequest struct {
	// raw transaction data to broadcast
	Hex string `json:"hex,omitempty"`
}

type BroadcastResponse struct {
	// the transaction hash
	Hash string `json:"hash,omitempty"`
}

type TransactionRequest struct {
	// tx hash
	Hash string `json:"hash,omitempty"`
}

type TransactionResponse struct {
	// the block hash
	BlockHash string `json:"block_hash,omitempty"`
	// the block number
	BlockNumber string `json:"block_number,omitempty"`
	// chain id
	ChainId string `json:"chain_id,omitempty"`
	// sent from
	FromAddress string `json:"from_address,omitempty"`
	// gas
	Gas string `json:"gas,omitempty"`
	// gas price
	GasPrice string `json:"gas_price,omitempty"`
	// tx hash
	Hash string `json:"hash,omitempty"`
	// input
	Input string `json:"input,omitempty"`
	// max fee per gas
	MaxFeePerGas string `json:"max_fee_per_gas,omitempty"`
	// max priority fee per gas
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas,omitempty"`
	// the nonce
	Nonce string `json:"nonce,omitempty"`
	R     string `json:"r,omitempty"`
	S     string `json:"s,omitempty"`
	// to address
	ToAddress string `json:"to_address,omitempty"`
	// transaction index
	TxIndex string `json:"tx_index,omitempty"`
	// type of transaction
	Type string `json:"type,omitempty"`
	V    string `json:"v,omitempty"`
	// value of transaction
	Value string `json:"value,omitempty"`
}
