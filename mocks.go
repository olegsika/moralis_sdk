package moralissdk

// NFTClientMock mock for moralis SDK
type NFTClientMock struct {
	GetNFTTransfersByBlockResult *GetNFTTransfersByBlockResponse
	GetNFTTransfersByBlockError  error

	GetWalletNFTsResult *GetWalletNFTsResponse
	GetWalletNFTsError  error

	GetWalletNFTTransfersResult *GetWalletNFTTransfersResponse
	GetWalletNFTTransfersError  error

	GetWalletNFTCollectionsResult *GetWalletNFTCollectionsResponse
	GetWalletNFTCollectionsError  error

	GetNFTTradesResult *GetNFTTradesResponse
	GetNFTTradesError  error

	GetNFTLowestPriceResult *GetNFTLowestPriceResponse
	GetNFTLowestPriceError  error

	SearchNFTsResult *SearchNFTsResponse
	SearchNFTsError  error

	GetNFTTransfersFromToBlockResult *GetNFTTransfersFromToBlockResponse
	GetNFTTransfersFromToBlockError  error

	GetContractNFTsResult *GetContractNFTsResponse
	GetContractNFTsError  error

	GetNFTContractTransfersResult *GetNFTContractTransfersResponse
	GetNFTContractTransfersError  error

	GetNFTOwnersResult *GetNFTOwnersResponse
	GetNFTOwnersError  error

	GetNFTContractMetadataResult *GetNFTContractMetadataResponse
	GetNFTContractMetadataError  error

	ReSyncMetadataResult *ReSyncMetadataResponse
	ReSyncMetadataError  error

	SyncNFTContractError error

	GetNFTMetadataResult *GetNFTMetadataResponse
	GetNFTMetadataError  error

	GetNFTTokenIDOwnersResult *GetNFTTokenIDOwnersResponse
	GetNFTTokenIDOwnersError  error

	GetNFTTransfersResult *GetNFTTransfersResponse
	GetNFTTransfersError  error
}

func (m *NFTClientMock) GetNFTTransfersByBlock(GetNFTTransfersByBlockParams, string) (*GetNFTTransfersByBlockResponse, error) {
	return m.GetNFTTransfersByBlockResult, m.GetNFTTransfersByBlockError
}

func (m *NFTClientMock) GetWalletNFTs(GetWalletNFTsParams, string) (*GetWalletNFTsResponse, error) {
	return m.GetWalletNFTsResult, m.GetWalletNFTsError
}

func (m *NFTClientMock) GetWalletNFTTransfers(GetWalletNFTTransfersParams, string) (*GetWalletNFTTransfersResponse, error) {
	return m.GetWalletNFTTransfersResult, m.GetWalletNFTTransfersError
}

func (m *NFTClientMock) GetWalletNFTCollections(GetWalletNFTCollectionsParams, string) (*GetWalletNFTCollectionsResponse, error) {
	return m.GetWalletNFTCollectionsResult, m.GetWalletNFTCollectionsError
}

func (m *NFTClientMock) GetNFTTrades(GetNFTTradesParams, string) (*GetNFTTradesResponse, error) {
	return m.GetNFTTradesResult, m.GetNFTTradesError
}

func (m *NFTClientMock) GetNFTLowestPrice(GetNFTLowestPriceParams, string) (*GetNFTLowestPriceResponse, error) {
	return m.GetNFTLowestPriceResult, m.GetNFTLowestPriceError
}

func (m *NFTClientMock) SearchNFTs(SearchNFTsParams) (*SearchNFTsResponse, error) {
	return m.SearchNFTsResult, m.SearchNFTsError
}

func (m *NFTClientMock) GetNFTTransfersFromToBlock(GetNFTTransfersFromToBlockParams) (*GetNFTTransfersFromToBlockResponse, error) {
	return m.GetNFTTransfersFromToBlockResult, m.GetNFTTransfersFromToBlockError
}

func (m *NFTClientMock) GetContractNFTs(GetContractNFTsParams, string) (*GetContractNFTsResponse, error) {
	return m.GetContractNFTsResult, m.GetContractNFTsError
}

func (m *NFTClientMock) GetNFTContractTransfers(GetNFTContractTransfersParams, string) (*GetNFTContractTransfersResponse, error) {
	return m.GetNFTContractTransfersResult, m.GetNFTContractTransfersError
}

func (m *NFTClientMock) GetNFTOwners(GetNFTOwnersParams, string) (*GetNFTOwnersResponse, error) {
	return m.GetNFTOwnersResult, m.GetNFTOwnersError
}

func (m *NFTClientMock) GetNFTContractMetadata(GetNFTContractMetadataParams, string) (*GetNFTContractMetadataResponse, error) {
	return m.GetNFTContractMetadataResult, m.GetNFTContractMetadataError
}

func (m *NFTClientMock) ReSyncMetadata(ReSyncMetadataParams, string, string) (*ReSyncMetadataResponse, error) {
	return m.ReSyncMetadataResult, m.ReSyncMetadataError
}

func (m *NFTClientMock) SyncNFTContract(SyncNFTContractParams, string) error {
	return m.SyncNFTContractError
}

func (m *NFTClientMock) GetNFTMetadata(GetNFTMetadataParams, string, string) (*GetNFTMetadataResponse, error) {
	return m.GetNFTMetadataResult, m.GetNFTMetadataError
}

func (m *NFTClientMock) GetNFTTokenIDOwners(GetNFTTokenIDOwnersParams, string, string) (*GetNFTTokenIDOwnersResponse, error) {
	return m.GetNFTTokenIDOwnersResult, m.GetNFTTokenIDOwnersError
}
func (m *NFTClientMock) GetNFTTransfers(GetNFTTransfersParams, string, string) (*GetNFTTransfersResponse, error) {
	return m.GetNFTTransfersResult, m.GetNFTTransfersError
}
