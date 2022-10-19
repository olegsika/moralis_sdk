package moralissdk

import (
	"fmt"
	"net/url"
	"testing"
	"time"
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
		validChain   = "sepolia"
		randomChain  = "test"
		defaultChain = "eth"

		validFormat   = "hex"
		randomFormad  = "test"
		defaultFormat = "decimal"

		validDirection   = "to"
		randomDirection  = "test"
		defaultDirection = "both"

		validMarketplace   = "opensea"
		randomMarketplace  = "test"
		defaultMarketplace = "opensea"

		validFilter   = "description"
		randomFilter  = "test"
		defaultFilter = "global"

		validFlag   = "metadata"
		randomFlag  = "test"
		defaultFlag = "uri"

		validMode   = "sync"
		randomMode  = "test"
		defaultMode = "async"

		fieldWithoutValidation = "field_without_validation"

		addressOne   = "address_one"
		addressTwo   = "address_two"
		addressThree = "address_three"

		tokenAddressOne   = "token_address_one"
		tokenAddressTwo   = "token_address_two"
		tokenAddressThree = "token_address_three"

		fieldWithoutValidationInt  = 1
		fieldWithoutValidationTime = time.Now()
	)

	type QueryParamsCommonStruct struct {
		Chain          string    `json:"chain,omitempty"`
		Format         string    `json:"format,omitempty"`
		Direction      string    `json:"direction,omitempty"`
		Marketplace    string    `json:"marketplace,omitempty"`
		Q              string    `json:"q,omitempty"`
		Filter         string    `json:"filter,omitempty"`
		Flag           string    `json:"flag,omitempty"`
		Mode           string    `json:"mode,omitempty"`
		FromBlock      int       `json:"from_block,omitempty"`
		ToBlock        int       `json:"to_block,omitempty"`
		FromDate       time.Time `json:"from_date,omitempty"`
		ToDate         time.Time `json:"to_date,omitempty"`
		Days           int       `json:"days,omitempty"`
		ProviderURL    string    `json:"provider_url,omitempty"`
		Addresses      []string  `json:"addresses,omitempty"`
		TokenAddresses []string  `json:"token_addresses,omitempty"`
		Subdomain      string    `json:"subdomain,omitempty"`
		Cursor         string    `json:"cursor,omitempty"`
		TotalRanges    int       `json:"total_ranges,omitempty"`
		Range          int       `json:"range,omitempty"`
		Limit          int       `json:"limit,omitempty"`
		Order          string    `json:"order,omitempty"`
	}

	tCases := []struct {
		name           string
		entryStruct    interface{}
		expectedResult string
		expectedErr    error
	}{
		{
			name: "Should be success. All fields valid",
			entryStruct: QueryParamsCommonStruct{
				Chain:          validChain,
				Cursor:         fieldWithoutValidation,
				Days:           fieldWithoutValidationInt,
				Direction:      validDirection,
				Filter:         validFilter,
				Flag:           validFlag,
				Format:         validFormat,
				FromBlock:      fieldWithoutValidationInt,
				FromDate:       fieldWithoutValidationTime,
				Limit:          fieldWithoutValidationInt,
				Marketplace:    validMarketplace,
				Mode:           validMode,
				Order:          fieldWithoutValidation,
				ProviderURL:    fieldWithoutValidation,
				Q:              fieldWithoutValidation,
				Range:          fieldWithoutValidationInt,
				Subdomain:      fieldWithoutValidation,
				ToBlock:        fieldWithoutValidationInt,
				ToDate:         fieldWithoutValidationTime,
				TotalRanges:    fieldWithoutValidationInt,
				TokenAddresses: []string{tokenAddressOne, tokenAddressTwo, tokenAddressThree},
				Addresses:      []string{addressOne, addressTwo, addressThree},
			},
			expectedResult: fmt.Sprintf("chain=%s&cursor=%s&days=%d&direction=%s&filter=%s&flag=%s&format=%s&from_block=%d&from_date=%v&limit=%d&marketplace=%s&mode=%s&order=%s&provider_url=%s&q=%s&range=%d&subdomain=%s&to_block=%d&to_date=%v&total_ranges=%d&token_addresses=%s&token_addresses=%s&token_addresses=%s&addresses=%s&addresses=%s&addresses=%s",
				validChain, fieldWithoutValidation, fieldWithoutValidationInt,
				validDirection, validFilter, validFlag, validFormat,
				fieldWithoutValidationInt, url.QueryEscape(fieldWithoutValidationTime.Format(time.RFC3339)),
				fieldWithoutValidationInt, validMarketplace, validMode,
				fieldWithoutValidation, fieldWithoutValidation,
				fieldWithoutValidation, fieldWithoutValidationInt, fieldWithoutValidation,
				fieldWithoutValidationInt, url.QueryEscape(fieldWithoutValidationTime.Format(time.RFC3339)),
				fieldWithoutValidationInt, tokenAddressOne, tokenAddressTwo, tokenAddressThree,
				addressOne, addressTwo, addressThree,
			),
		},
		{
			name: "Chain Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Chain: randomChain,
			},
			expectedResult: fmt.Sprintf("chain=%s", defaultChain),
		},
		{
			name: "Format Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Format: randomFormad,
			},
			expectedResult: fmt.Sprintf("format=%s", defaultFormat),
		},
		{
			name: "Direction Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Direction: randomDirection,
			},
			expectedResult: fmt.Sprintf("direction=%s", defaultDirection),
		},
		{
			name: "Marketplace Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Marketplace: randomMarketplace,
			},
			expectedResult: fmt.Sprintf("marketplace=%s", defaultMarketplace),
		},
		{
			name: "Filter Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Filter: randomFilter,
			},
			expectedResult: fmt.Sprintf("filter=%s", defaultFilter),
		},
		{
			name: "Flag Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Flag: randomFlag,
			},
			expectedResult: fmt.Sprintf("flag=%s", defaultFlag),
		},
		{
			name: "Mode Field. Invalid",
			entryStruct: QueryParamsCommonStruct{
				Mode: randomMode,
			},
			expectedResult: fmt.Sprintf("mode=%s", defaultMode),
		},
		{
			name: "Addresses Field Nil",
			entryStruct: QueryParamsCommonStruct{
				Addresses: nil,
			},
			expectedResult: "",
		},
		{
			name: "Token Addresses Field",
			entryStruct: QueryParamsCommonStruct{
				TokenAddresses: nil,
			},
			expectedResult: "",
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
