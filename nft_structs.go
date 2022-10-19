package moralissdk

import "time"

type GetNFTTransfersByBlockParams struct {
	Chain     string `json:"chain"`
	Subdomain string `json:"subdomain,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
type GetNFTTransfersByBlockResponse struct {
	BaseResponse
	Cursor string                                 `json:"cursor"`
	Result []GetNFTTransfersByBlockResponseResult `json:"result"`
	BlockExistsIndexCompleteFields
}
type GetNFTTransfersByBlockResponseResult struct {
	TokenAddress     string `json:"token_address"`
	TokenID          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Value            string `json:"value"`
	Amount           string `json:"amount"`
	ContractType     string `json:"contract_type"`
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionType  string `json:"transaction_type"`
	TransactionIndex int    `json:"transaction_index"`
	LogIndex         int    `json:"log_index"`
	Operator         string `json:"operator"`
}

type GetWalletNFTsParams struct {
	Chain          string   `json:"chain"`
	Format         string   `json:"format,omitempty"`
	TokenAddresses []string `json:"token_addresses,omitempty"`
	Cursor         string   `json:"cursor,omitempty"`
	Limit          int      `json:"limit,omitempty"`
}
type GetWalletNFTsResponse struct {
	BaseResponse
	Status string `json:"status"`
	Cursor string `json:"cursor"`
	Result []GetWalletNFTsResponseResult
}
type GetWalletNFTsResponseResult struct {
	TokenAddress      string    `json:"token_address"`
	TokenID           string    `json:"token_id"`
	ContractType      string    `json:"contract_type"`
	OwnerOf           string    `json:"owner_of"`
	BlockNumber       string    `json:"block_number"`
	BlockNumberMinted string    `json:"block_number_minted"`
	TokenURI          string    `json:"token_uri"`
	Metadata          string    `json:"metadata"`
	Amount            string    `json:"amount"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	TokenHash         string    `json:"token_hash"`
	LastTokenURISync  time.Time `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time `json:"last_metadata_sync"`
}

type GetWalletNFTTransfersParams struct {
	Chain     string `json:"chain"`
	Format    string `json:"format"`
	Direction string `json:"direction"`
	Cursor    string `json:"cursor,omitempty"`
	FromBlock int    `json:"from_block,omitempty"`
	ToBlock   int    `json:"to_block,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
type GetWalletNFTTransfersResponse struct {
	BaseResponse
	Cursor string                                `json:"cursor"`
	Result []GetWalletNFTTransfersResponseResult `json:"result"`
	BlockExistsIndexCompleteFields
}
type GetWalletNFTTransfersResponseResult struct {
	TokenAddress     string `json:"token_address"`
	TokenID          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Value            string `json:"value"`
	Amount           string `json:"amount"`
	ContractType     string `json:"contract_type"`
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionType  string `json:"transaction_type"`
	TransactionIndex int    `json:"transaction_index"`
	LogIndex         int    `json:"log_index"`
	Operator         string `json:"operator"`
}

type GetWalletNFTCollectionsParams struct {
	Chain  string `json:"chain"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
type GetWalletNFTCollectionsResponse struct {
	BaseResponse
	Status string                                  `json:"status"`
	Cursor string                                  `json:"cursor"`
	Result []GetWalletNFTCollectionsResponseResult `json:"result"`
}
type GetWalletNFTCollectionsResponseResult struct {
	TokenAddress string `json:"token_address"`
	ContractType string `json:"contract_type"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
}

type GetNFTTradesParams struct {
	Chain       string    `json:"chain"`
	FromBlock   int       `json:"from_block,omitempty"`
	ToBlock     int       `json:"to_block,omitempty"`
	FromDate    time.Time `json:"from_date,omitempty"`
	ToDate      time.Time `json:"to_date,omitempty"`
	ProviderURL string    `json:"provider_url,omitempty"`
	Marketplace string    `json:"marketplace"`
	Cursor      string    `json:"cursor,omitempty"`
	Limit       int       `json:"limit,omitempty"`
}
type GetNFTTradesResponse struct {
	BaseResponse
	Result []GetNFTTradesResponseResult `json:"result"`
}
type GetNFTTradesResponseResult struct {
	TransactionHash    string   `json:"transaction_hash"`
	TransactionIndex   string   `json:"transaction_index"`
	TokenIds           []string `json:"token_ids"`
	SellerAddress      string   `json:"seller_address"`
	BuyerAddress       string   `json:"buyer_address"`
	MarketplaceAddress string   `json:"marketplace_address"`
	Price              string   `json:"price"`
	BlockTimestamp     string   `json:"block_timestamp"`
	BlockNumber        string   `json:"block_number"`
	BlockHash          string   `json:"block_hash"`
}

type GetNFTLowestPriceParams struct {
	Chain       string `json:"chain"`
	Days        int    `json:"days,omitempty"`
	ProviderURL string `json:"provider_url,omitempty"`
	Marketplace string `json:"marketplace"`
}
type GetNFTLowestPriceResponse struct {
	TransactionHash    string   `json:"transaction_hash"`
	TransactionIndex   string   `json:"transaction_index"`
	TokenIds           []string `json:"token_ids"`
	SellerAddress      string   `json:"seller_address"`
	BuyerAddress       string   `json:"buyer_address"`
	MarketplaceAddress string   `json:"marketplace_address"`
	Price              string   `json:"price"`
	BlockTimestamp     string   `json:"block_timestamp"`
	BlockNumber        string   `json:"block_number"`
	BlockHash          string   `json:"block_hash"`
}

type SearchNFTsParams struct {
	Chain     string    `json:"chain"`
	Format    string    `json:"format"`
	Q         string    `json:"q"`
	Filter    string    `json:"filter"`
	FromBlock int       `json:"from_block,omitempty"`
	ToBlock   int       `json:"to_block,omitempty"`
	FromDate  time.Time `json:"from_date,omitempty"`
	ToDate    time.Time `json:"to_date,omitempty"`
	Addresses []string  `json:"addresses,omitempty"`
	Cursor    string    `json:"cursor,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}
type SearchNFTsResponse struct {
	BaseResponse
	Result []SearchNFTsResponseResult `json:"result"`
}
type SearchNFTsResponseResult struct {
	TokenId             string      `json:"token_id"`
	TokenAddress        string      `json:"token_address"`
	TokenUri            string      `json:"token_uri"`
	Metadata            string      `json:"metadata"`
	IsValid             int         `json:"is_valid"`
	Syncing             int         `json:"syncing"`
	Frozen              int         `json:"frozen"`
	Resyncing           int         `json:"resyncing"`
	ContractType        string      `json:"contract_type"`
	TokenHash           string      `json:"token_hash"`
	BatchId             string      `json:"batch_id"`
	MetadataName        string      `json:"metadata_name"`
	MetadataDescription string      `json:"metadata_description"`
	MetadataAttributes  string      `json:"metadata_attributes"`
	BlockNumberMinted   string      `json:"block_number_minted"`
	OpenseaLookup       interface{} `json:"opensea_lookup"`
	MinterAddress       string      `json:"minter_address"`
	TransactionMinted   string      `json:"transaction_minted"`
	FrozenLogIndex      interface{} `json:"frozen_log_index"`
	Imported            interface{} `json:"imported"`
	LastTokenUriSync    time.Time   `json:"last_token_uri_sync"`
	LastMetadataSync    time.Time   `json:"last_metadata_sync"`
	CreatedAt           time.Time   `json:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt"`
}

type GetNFTTransfersFromToBlockParams struct {
	Chain     string    `json:"chain"`
	FromBlock int       `json:"from_block,omitempty"`
	ToBlock   int       `json:"to_block,omitempty"`
	FromDate  time.Time `json:"from_date,omitempty"`
	ToDate    time.Time `json:"to_date,omitempty"`
	Format    string    `json:"format"`
	Limit     int       `json:"limit,omitempty"`
	Cursor    string    `json:"cursor,omitempty"`
}
type GetNFTTransfersFromToBlockResponse struct {
	BaseResponse
	Cursor string                                     `json:"cursor"`
	Result []GetNFTTransfersFromToBlockResponseResult `json:"result"`
	BlockExistsIndexCompleteFields
}
type GetNFTTransfersFromToBlockResponseResult struct {
	TokenAddress     string `json:"token_address"`
	TokenId          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Value            string `json:"value"`
	Amount           string `json:"amount"`
	ContractType     string `json:"contract_type"`
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionType  string `json:"transaction_type"`
	TransactionIndex int    `json:"transaction_index"`
	LogIndex         int    `json:"log_index"`
	Operator         string `json:"operator"`
}

type GetContractNFTsParams struct {
	Chain       string `json:"chain"`
	Format      string `json:"format"`
	Limit       int    `json:"limit,omitempty"`
	TotalRanges int    `json:"total_ranges,omitempty"`
	Range       int    `json:"range,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
}
type GetContractNFTsResponse struct {
	BaseResponse
	Cursor string                          `json:"cursor"`
	Result []GetContractNFTsResponseResult `json:"result"`
}
type GetContractNFTsResponseResult struct {
	TokenAddress      string `json:"token_address"`
	TokenId           string `json:"token_id"`
	OwnerOf           string `json:"owner_of"`
	TokenHash         string `json:"token_hash"`
	BlockNumber       string `json:"block_number"`
	BlockNumberMinted string `json:"block_number_minted"`
	ContractType      string `json:"contract_type"`
	TokenUri          string `json:"token_uri"`
	Metadata          string `json:"metadata"`
	LastTokenUriSync  string `json:"last_token_uri_sync"`
	LastMetadataSync  string `json:"last_metadata_sync"`
	Amount            string `json:"amount"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
}

type GetNFTContractTransfersParams struct {
	Chain  string `json:"chain"`
	Format string `json:"format"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type GetNFTContractTransfersResponse struct {
	BaseResponse
	Cursor string                                  `json:"cursor"`
	Result []GetNFTContractTransfersResponseResult `json:"result"`
	BlockExistsIndexCompleteFields
}
type GetNFTContractTransfersResponseResult struct {
	TokenAddress     string `json:"token_address"`
	TokenId          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Value            string `json:"value"`
	Amount           string `json:"amount"`
	ContractType     string `json:"contract_type"`
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionType  string `json:"transaction_type"`
	TransactionIndex int    `json:"transaction_index"`
	LogIndex         int    `json:"log_index"`
	Operator         string `json:"operator"`
}

type GetNFTOwnersParams struct {
	Chain  string `json:"chain"`
	Format string `json:"format"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type GetNFTOwnersResponse struct {
	BaseResponse
	Status string                       `json:"status"`
	Cursor string                       `json:"cursor"`
	Result []GetNFTOwnersResponseResult `json:"result"`
}
type GetNFTOwnersResponseResult struct {
	TokenAddress      string    `json:"token_address"`
	TokenId           string    `json:"token_id"`
	ContractType      string    `json:"contract_type"`
	OwnerOf           string    `json:"owner_of"`
	BlockNumber       string    `json:"block_number"`
	BlockNumberMinted string    `json:"block_number_minted"`
	TokenUri          string    `json:"token_uri"`
	Metadata          string    `json:"metadata"`
	Amount            string    `json:"amount"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	TokenHash         string    `json:"token_hash"`
	LastTokenUriSync  time.Time `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time `json:"last_metadata_sync"`
}

type GetNFTContractMetadataParams struct {
	Chain string `json:"chain"`
}
type GetNFTContractMetadataResponse struct {
	TokenAddress string `json:"token_address"`
	Name         string `json:"name"`
	SyncedAt     string `json:"synced_at"`
	Symbol       string `json:"symbol"`
	ContractType string `json:"contract_type"`
}

type ReSyncMetadataParams struct {
	Chain string `json:"chain"`
	Flag  string `json:"flag"`
	Mode  string `json:"mode"`
}
type ReSyncMetadataResponse struct {
	Status string `json:"status"`
}

type SyncNFTContractParams struct {
	Chain string `json:"chain"`
}

type GetNFTMetadataParams struct {
	Chain  string `json:"chain"`
	Format string `json:"format"`
}
type GetNFTMetadataResponse struct {
	TokenAddress      string `json:"token_address"`
	TokenId           string `json:"token_id"`
	OwnerOf           string `json:"owner_of"`
	TokenHash         string `json:"token_hash"`
	BlockNumber       string `json:"block_number"`
	BlockNumberMinted string `json:"block_number_minted"`
	ContractType      string `json:"contract_type"`
	TokenUri          string `json:"token_uri"`
	Metadata          string `json:"metadata"`
	LastTokenUriSync  string `json:"last_token_uri_sync"`
	LastMetadataSync  string `json:"last_metadata_sync"`
	Amount            string `json:"amount"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
}

type GetNFTTokenIDOwnersParams struct {
	Chain  string `json:"chain"`
	Format string `json:"format"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type GetNFTTokenIDOwnersResponse struct {
	BaseResponse
	Status string                              `json:"status"`
	Cursor string                              `json:"cursor"`
	Result []GetNFTTokenIDOwnersResponseResult `json:"result"`
}
type GetNFTTokenIDOwnersResponseResult struct {
	TokenAddress      string    `json:"token_address"`
	TokenId           string    `json:"token_id"`
	ContractType      string    `json:"contract_type"`
	OwnerOf           string    `json:"owner_of"`
	BlockNumber       string    `json:"block_number"`
	BlockNumberMinted string    `json:"block_number_minted"`
	TokenUri          string    `json:"token_uri"`
	Metadata          string    `json:"metadata"`
	Amount            string    `json:"amount"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	TokenHash         string    `json:"token_hash"`
	LastTokenUriSync  time.Time `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time `json:"last_metadata_sync"`
}

type GetNFTTransfersParams struct {
	Chain  string `json:"chain"`
	Format string `json:"format"`
	Limit  int    `json:"limit,omitempty"`
	Order  string `json:"order,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type GetNFTTransfersResponse struct {
	BaseResponse
	Cursor string                          `json:"cursor"`
	Result []GetNFTTransfersResponseResult `json:"result"`
	BlockExistsIndexCompleteFields
}
type GetNFTTransfersResponseResult struct {
	TokenAddress     string `json:"token_address"`
	TokenId          string `json:"token_id"`
	FromAddress      string `json:"from_address"`
	ToAddress        string `json:"to_address"`
	Value            string `json:"value"`
	Amount           string `json:"amount"`
	ContractType     string `json:"contract_type"`
	BlockNumber      string `json:"block_number"`
	BlockTimestamp   string `json:"block_timestamp"`
	BlockHash        string `json:"block_hash"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionType  string `json:"transaction_type"`
	TransactionIndex int    `json:"transaction_index"`
	LogIndex         int    `json:"log_index"`
	Operator         string `json:"operator"`
}

type BaseResponse struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type BlockExistsIndexCompleteFields struct {
	BlockExists   bool `json:"block_exists"`
	IndexComplete bool `json:"index_complete"`
}
