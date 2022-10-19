package moralissdk

type Client interface {
	GetNFTTransfersByBlock(GetNFTTransfersByBlockParams, string) (*GetNFTTransfersByBlockResponse, error)
	GetWalletNFTs(GetWalletNFTsParams, string) (*GetWalletNFTsResponse, error)
	GetWalletNFTTransfers(GetWalletNFTTransfersParams, string) (*GetWalletNFTTransfersResponse, error)
	GetWalletNFTCollections(GetWalletNFTCollectionsParams, string) (*GetWalletNFTCollectionsResponse, error)
	GetNFTTrades(GetNFTTradesParams, string) (*GetNFTTradesResponse, error)
	GetNFTLowestPrice(GetNFTLowestPriceParams, string) (*GetNFTLowestPriceResponse, error)
	SearchNFTs(SearchNFTsParams) (*SearchNFTsResponse, error)
	GetNFTTransfersFromToBlock(GetNFTTransfersFromToBlockParams) (*GetNFTTransfersFromToBlockResponse, error)
	GetContractNFTs(GetContractNFTsParams, string) (*GetContractNFTsResponse, error)
	GetNFTContractTransfers(GetNFTContractTransfersParams, string) (*GetNFTContractTransfersResponse, error)
	GetNFTOwners(GetNFTOwnersParams, string) (*GetNFTOwnersResponse, error)
	GetNFTContractMetadata(GetNFTContractMetadataParams, string) (*GetNFTContractMetadataResponse, error)
	ReSyncMetadata(ReSyncMetadataParams, string, string) (*ReSyncMetadataResponse, error)
	SyncNFTContract(SyncNFTContractParams, string) error
	GetNFTMetadata(GetNFTMetadataParams, string, string) (*GetNFTMetadataResponse, error)
	GetNFTTokenIDOwners(GetNFTTokenIDOwnersParams, string, string) (*GetNFTTokenIDOwnersResponse, error)
	GetNFTTransfers(GetNFTTransfersParams, string, string) (*GetNFTTransfersResponse, error)
}
