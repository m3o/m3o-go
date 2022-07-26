package wallet

import (
	"go.m3o.com/client"
)

type Wallet interface {
	Balance(*BalanceRequest) (*BalanceResponse, error)
	Create(*CreateRequest) (*CreateResponse, error)
	Credit(*CreditRequest) (*CreditResponse, error)
	Debit(*DebitRequest) (*DebitResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Transactions(*TransactionsRequest) (*TransactionsResponse, error)
	Transfer(*TransferRequest) (*TransferResponse, error)
}

func NewWalletService(token string) *WalletService {
	return &WalletService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type WalletService struct {
	client *client.Client
}

// Get the balance of a wallet
func (t *WalletService) Balance(request *BalanceRequest) (*BalanceResponse, error) {

	rsp := &BalanceResponse{}
	return rsp, t.client.Call("wallet", "Balance", request, rsp)

}

// Create a new wallet
func (t *WalletService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("wallet", "Create", request, rsp)

}

// Add credit to a wallet
func (t *WalletService) Credit(request *CreditRequest) (*CreditResponse, error) {

	rsp := &CreditResponse{}
	return rsp, t.client.Call("wallet", "Credit", request, rsp)

}

// Debit a wallet
func (t *WalletService) Debit(request *DebitRequest) (*DebitResponse, error) {

	rsp := &DebitResponse{}
	return rsp, t.client.Call("wallet", "Debit", request, rsp)

}

// Delete a wallet
func (t *WalletService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("wallet", "Delete", request, rsp)

}

// List your wallets
func (t *WalletService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("wallet", "List", request, rsp)

}

// Get wallet by id
func (t *WalletService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("wallet", "Read", request, rsp)

}

// List the transactions for a wallet
func (t *WalletService) Transactions(request *TransactionsRequest) (*TransactionsResponse, error) {

	rsp := &TransactionsResponse{}
	return rsp, t.client.Call("wallet", "Transactions", request, rsp)

}

// Make a transfer from one wallet to another
func (t *WalletService) Transfer(request *TransferRequest) (*TransferResponse, error) {

	rsp := &TransferResponse{}
	return rsp, t.client.Call("wallet", "Transfer", request, rsp)

}

type Account struct {
	// current balance
	Balance int64 `json:"balance,string,omitempty"`
	// currency of balance
	Currency string `json:"currency,omitempty"`
	// description of the wallet
	Description string `json:"description,omitempty"`
	// wallet id
	Id string `json:"id,omitempty"`
	// name of the wallet
	Name string `json:"name,omitempty"`
}

type BalanceRequest struct {
	// wallet id
	Id string `json:"id,omitempty"`
}

type BalanceResponse struct {
	// current balance
	Balance int64 `json:"balance,string,omitempty"`
}

type CreateRequest struct {
	// specify a currency
	Currency string `json:"currency,omitempty"`
	// description for wallet
	Description string `json:"description,omitempty"`
	// optional id
	Id string `json:"id,omitempty"`
	// name of the wallet
	Name string `json:"name,omitempty"`
}

type CreateResponse struct {
	// the wallet created
	Account *Account `json:"account,omitempty"`
}

type CreditRequest struct {
	// amount to credit
	Amount int64 `json:"amount,string,omitempty"`
	// wallet id
	Id string `json:"id,omitempty"`
	// idempotency key
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// reference note
	Reference string `json:"reference,omitempty"`
	// if the transaction is visible
	Visible bool `json:"visible,omitempty"`
}

type CreditResponse struct {
	// the new balance
	Balance int64 `json:"balance,string,omitempty"`
}

type DebitRequest struct {
	// amount to debit
	Amount int64 `json:"amount,string,omitempty"`
	// wallet
	Id string `json:"id,omitempty"`
	// idempotency key
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// reference note
	Reference string `json:"reference,omitempty"`
	// if the transaction is visible
	Visible bool `json:"visible,omitempty"`
}

type DebitResponse struct {
	// the new balance
	Balance int64 `json:"balance,string,omitempty"`
}

type DeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type DeleteResponse struct {
}

type ListRequest struct {
}

type ListResponse struct {
	Accounts []Account `json:"accounts,omitempty"`
}

type ReadRequest struct {
	// wallet id
	Id string `json:"id,omitempty"`
}

type ReadResponse struct {
	Account *Account `json:"account,omitempty"`
}

type Transaction struct {
	// amount in transaction
	Amount int64 `json:"amount,string,omitempty"`
	// time of transaction
	Created string `json:"created,omitempty"`
	// unique id of transaction
	Id string `json:"id,omitempty"`
	// associated metadata
	Metadata map[string]string `json:"metadata,omitempty"`
	// reference note
	Reference string `json:"reference,omitempty"`
}

type TransactionsRequest struct {
	// wallet id
	Id string `json:"id,omitempty"`
}

type TransactionsResponse struct {
	// list of transactions
	Transactions []Transaction `json:"transactions,omitempty"`
}

type TransferRequest struct {
	// amount to transfer
	Amount int64 `json:"amount,string,omitempty"`
	// from wallet id
	FromId string `json:"from_id,omitempty"`
	// reference
	Reference string `json:"reference,omitempty"`
	// to wallet id
	ToId string `json:"to_id,omitempty"`
	// visible?
	Visible bool `json:"visible,omitempty"`
}

type TransferResponse struct {
}
