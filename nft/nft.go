package nft

import (
	"go.m3o.com/client"
)

type Nft interface {
	Asset(*AssetRequest) (*AssetResponse, error)
	Assets(*AssetsRequest) (*AssetsResponse, error)
	Collection(*CollectionRequest) (*CollectionResponse, error)
	Collections(*CollectionsRequest) (*CollectionsResponse, error)
	Create(*CreateRequest) (*CreateResponse, error)
}

func NewNftService(token string) *NftService {
	return &NftService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type NftService struct {
	client *client.Client
}

// Get a single asset by the contract
func (t *NftService) Asset(request *AssetRequest) (*AssetResponse, error) {

	rsp := &AssetResponse{}
	return rsp, t.client.Call("nft", "Asset", request, rsp)

}

// Return a list of assets
func (t *NftService) Assets(request *AssetsRequest) (*AssetsResponse, error) {

	rsp := &AssetsResponse{}
	return rsp, t.client.Call("nft", "Assets", request, rsp)

}

// Get a collection by its slug
func (t *NftService) Collection(request *CollectionRequest) (*CollectionResponse, error) {

	rsp := &CollectionResponse{}
	return rsp, t.client.Call("nft", "Collection", request, rsp)

}

// Get a list of collections
func (t *NftService) Collections(request *CollectionsRequest) (*CollectionsResponse, error) {

	rsp := &CollectionsResponse{}
	return rsp, t.client.Call("nft", "Collections", request, rsp)

}

// Create your own NFT (coming soon)
func (t *NftService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("nft", "Create", request, rsp)

}

type Asset struct {
	// associated collection
	Collection *Collection `json:"collection"`
	// asset contract
	Contract *Contract `json:"contract"`
	// Creator of the NFT
	Creator *User `json:"creator"`
	// related description
	Description string `json:"description"`
	// id of the asset
	Id int32 `json:"id"`
	// the image url
	ImageUrl string `json:"image_url"`
	// last time sold
	LastSale *Sale `json:"last_sale"`
	// listing date
	ListingDate string `json:"listing_date"`
	// name of the asset
	Name string `json:"name"`
	// Owner of the NFT
	Owner *User `json:"owner"`
	// the permalink
	Permalink string `json:"permalink"`
	// is it a presale
	Presale bool `json:"presale"`
	// number of sales
	Sales int32 `json:"sales"`
	// the token id
	TokenId string `json:"token_id"`
	// traits associated with the item
	Traits []map[string]interface{} `json:"traits"`
}

type AssetRequest struct {
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
}

type AssetResponse struct {
	Asset *Asset `json:"asset"`
}

type AssetsRequest struct {
	// limit to members of a collection by slug name (case sensitive)
	Collection string `json:"collection"`
	// A cursor pointing to the page to retrieve
	Cursor string `json:"cursor"`
	// limit returned assets
	Limit int32 `json:"limit"`
	// DEPRECATED offset for pagination, please use cursor instead
	Offset int32 `json:"offset"`
	// order "asc" or "desc"
	Order string `json:"order"`
	// order by "sale_date", "sale_count", "sale_price", "total_price"
	OrderBy string `json:"order_by"`
}

type AssetsResponse struct {
	// list of assets
	Assets []Asset `json:"assets"`
	// A cursor to be supplied to retrieve the next page of results
	Next string `json:"next"`
	// A cursor to be supplied to retrieve the previous page of results
	Previous string `json:"previous"`
}

type Collection struct {
	// image used in the banner for the collection
	BannerImageUrl string `json:"banner_image_url"`
	// creation time
	CreatedAt string `json:"created_at"`
	// description of the collection
	Description string `json:"description"`
	// approved editors for this collection
	Editors []string `json:"editors"`
	// external link to the original website for the collection
	ExternalLink string `json:"external_link"`
	// an image for the collection
	ImageUrl string `json:"image_url"`
	// name of the collection
	Name string `json:"name"`
	// the payment tokens accepted for this collection
	PaymentTokens *Token `json:"payment_tokens"`
	// payout address for the collection's royalties
	PayoutAddress string `json:"payout_address"`
	// a list of the contracts associated with this collection
	PrimaryAssetContracts *Contract `json:"primary_asset_contracts"`
	// the collection's approval status on OpenSea
	SafelistRequestStatus string `json:"safelist_request_status"`
	// the fees that get paid out when a sale is made
	SellerFees string `json:"seller_fees"`
	// collection slug
	Slug string `json:"slug"`
	// sales statistics associated with the collection
	Stats map[string]interface{} `json:"stats"`
	// listing of all the trait types available within this collection
	Traits map[string]interface{} `json:"traits"`
}

type CollectionRequest struct {
	Slug string `json:"slug"`
}

type CollectionResponse struct {
	Collection *Collection `json:"collection"`
}

type CollectionsRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}

type Contract struct {
	// ethereum address
	Address string `json:"address"`
	// timestamp of creation
	CreatedAt string `json:"created_at"`
	// description of contract
	Description string `json:"description"`
	// name of contract
	Name string `json:"name"`
	// owner id
	Owner int32 `json:"owner"`
	// payout address
	PayoutAddress string `json:"payout_address"`
	// aka "ERC1155"
	Schema string `json:"schema"`
	// seller fees
	SellerFees string `json:"seller_fees"`
	// related symbol
	Symbol string `json:"symbol"`
	// type of contract e.g "semi-fungible"
	Type string `json:"type"`
}

type CreateRequest struct {
	// data if not image
	Data string `json:"data"`
	// description
	Description string `json:"description"`
	// image data
	Image string `json:"image"`
	// name of the NFT
	Name string `json:"name"`
}

type CreateResponse struct {
	Asset *Asset `json:"asset"`
}

type Sale struct {
	AssetDecimals  int32        `json:"asset_decimals"`
	AssetTokenId   string       `json:"asset_token_id"`
	CreatedAt      string       `json:"created_at"`
	EventTimestamp string       `json:"event_timestamp"`
	EventType      string       `json:"event_type"`
	PaymentToken   *Token       `json:"payment_token"`
	Quantity       string       `json:"quantity"`
	TotalPrice     string       `json:"total_price"`
	Transaction    *Transaction `json:"transaction"`
}

type Token struct {
	Address  string `json:"address"`
	Decimals int32  `json:"decimals"`
	EthPrice string `json:"eth_price"`
	Id       int32  `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	UsdPrice string `json:"usd_price"`
}

type Transaction struct {
	BlockHash        string `json:"block_hash"`
	BlockNumber      string `json:"block_number"`
	FromAccount      *User  `json:"from_account"`
	Id               int32  `json:"id"`
	Timestamp        string `json:"timestamp"`
	ToAccount        *User  `json:"to_account"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionIndex string `json:"transaction_index"`
}

type User struct {
	Address    string `json:"address"`
	ProfileUrl string `json:"profile_url"`
	Username   string `json:"username"`
}
