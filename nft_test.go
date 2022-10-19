package moralissdk

import (
	"testing"
)

func TestGetNFTTransfersByBlock(t *testing.T) {

}

func TestGetWalletNFTs(t *testing.T) {

}
func TestGetWalletNFTTransfers(t *testing.T) {

}
func TestGetWalletNFTCollections(t *testing.T) {

}
func TestGetNFTTrades(t *testing.T) {

}
func TestGetNFTLowestPrice(t *testing.T) {

}
func TestSearchNFTs(t *testing.T) {

}

func TestGetNFTTransfersFromToBlock(t *testing.T) {

}
func TestGetContractNFTs(t *testing.T) {

}
func TestGetNFTContractTransfers(t *testing.T) {

}
func TestGetNFTOwners(t *testing.T) {

}
func TestGetNFTContractMetadata(t *testing.T) {

}
func TestReSyncMetadata(t *testing.T) {

}
func TestSyncNFTContract(t *testing.T) {

}
func TestGetNFTMetadata(t *testing.T) {

}
func TestGetNFTTokenIDOwners(t *testing.T) {

}
func TestGetNFTTransfers(t *testing.T) {

}

func TestBuildQueryParams(t *testing.T) {
	var (
		validChain     = "sepolia"
		randomChain    = "test"
		validSubdomain = "test_subdomain"
		validCursor    = "test_cursor"
		validLimit     = 10
	)

	tCases := []struct {
		name           string
		entryStruct    interface{}
		expectedResult string
		expectedErr    error
	}{
		{
			name: "GetNFTTransfersByBlockParams. All fields Valid",
			entryStruct: GetNFTTransfersByBlockParams{
				Chain:     validChain,
				Subdomain: validSubdomain,
				Cursor:    validCursor,
				Limit:     validLimit,
			},
			expectedResult: "chain=sepolia&cursor=test_cursor&limit=10&subdomain=test_subdomain",
		},
		{
			name: "GetNFTTransfersByBlockParams. All fields Chain Invalid",
			entryStruct: GetNFTTransfersByBlockParams{
				Chain:     randomChain,
				Subdomain: validSubdomain,
				Cursor:    validCursor,
				Limit:     validLimit,
			},
			expectedResult: "chain=eth&cursor=test_cursor&limit=10&subdomain=test_subdomain",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			res, err := BuildQueryParams(tc.entryStruct)
			if !IsErrorEqual(tc.expectedErr, err) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			if res != tc.expectedResult {
				tt.Errorf("expected result = %s, but actual result %s", tc.expectedResult, res)
				return
			}
		})
	}
}
