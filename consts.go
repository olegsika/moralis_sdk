package moralissdk

const (
	// NFTBaseURL Base URL
	NFTBaseURL = "https://deep-index.moralis.io/api/v2/%s"

	// APIKeyHeader Header name
	APIKeyHeader = "X-API-Key"
)

const (
	ChainField          = "chain"
	FormatField         = "format"
	DirectionField      = "direction"
	MarketPlaceField    = "marketplace"
	FilterField         = "filter"
	FlagField           = "flag"
	ModeField           = "mode"
	FromDateField       = "from_date"
	ToDateField         = "to_date"
	TokenAddressesField = "token_addresses"
	AdressesField       = "addresses"
)

const (
	GetNFTTransfersByBlockURL     = "block/%s/nft/transfers?%s"
	GetWalletNFTsURL              = "%s/nft?%s"
	GetWalletNFTTransfersURL      = "%s/nft/transfers?%s"
	GetWalletNFTCollectionsURL    = "%s/nft/collections?%s"
	GetNFTTradesURL               = "nft/%s/trades?%s"
	GetNFTLowestPriceURL          = "nft/%s/lowestprice?%s"
	SearchNFTsURL                 = "nft/search?%s"
	GetNFTTransfersFromToBlockURL = "nft/transfers?%s"
	GetContractNFTsURL            = "nft/%s?%s"
	GetNFTContractTransfersURL    = "nft/%s/transfers?%s"
	GetNFTOwnersURL               = "nft/%s/owners?%s"
	GetNFTContractMetadataURL     = "nft/%s/metadata?%s"
	ReSyncMetadataURL             = "nft/%s/%s/metadata/resync?%s"
	SyncNFTContractURL            = "nft/%s/sync?%s"
	GetNFTMetadataURL             = "nft/%s/%s?%s"
	GetNFTTokenIDOwnersURL        = "nft/%s/%s/owners?%s"
	GetNFTTransfersURL            = "nft/%s/%s/transfers?%s"
)
