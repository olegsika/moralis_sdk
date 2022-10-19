package moralissdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client manages requests to the Moralis API
type client struct {
	// Moralis ApiKey
	ApiKey string
}

func NewClient(apiKey string) Client {
	return &client{
		ApiKey: apiKey,
	}
}

func (c *client) GetNFTTransfersByBlock(params GetNFTTransfersByBlockParams, blockNumberOrHash string) (*GetNFTTransfersByBlockResponse, error) {
	if err := ValidateBlockNumberOrHash(blockNumberOrHash); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTTransfersByBlockURL, blockNumberOrHash, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetNFTTransfersByBlockResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetWalletNFTs(params GetWalletNFTsParams, address string) (*GetWalletNFTsResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetWalletNFTsURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetWalletNFTsResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetWalletNFTTransfers(params GetWalletNFTTransfersParams, address string) (*GetWalletNFTTransfersResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetWalletNFTTransfersURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetWalletNFTTransfersResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetWalletNFTCollections(params GetWalletNFTCollectionsParams, address string) (*GetWalletNFTCollectionsResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetWalletNFTCollectionsURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetWalletNFTCollectionsResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTTrades(params GetNFTTradesParams, address string) (*GetNFTTradesResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTTradesURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetNFTTradesResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTLowestPrice(params GetNFTLowestPriceParams, address string) (*GetNFTLowestPriceResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTLowestPriceURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetNFTLowestPriceResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) SearchNFTs(params SearchNFTsParams) (*SearchNFTsResponse, error) {
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(SearchNFTsURL, queryParams))
	if err != nil {
		return nil, err
	}
	res := &SearchNFTsResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTTransfersFromToBlock(params GetNFTTransfersFromToBlockParams) (*GetNFTTransfersFromToBlockResponse, error) {
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTTransfersFromToBlockURL, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetNFTTransfersFromToBlockResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetContractNFTs(params GetContractNFTsParams, address string) (*GetContractNFTsResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetContractNFTsURL, address, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetContractNFTsResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTContractTransfers(params GetNFTContractTransfersParams, address string) (*GetNFTContractTransfersResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}

	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTContractTransfersURL, address, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetNFTContractTransfersResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetNFTOwners(params GetNFTOwnersParams, address string) (*GetNFTOwnersResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTOwnersURL, address, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetNFTOwnersResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTContractMetadata(params GetNFTContractMetadataParams, address string) (*GetNFTContractMetadataResponse, error) {
	if err := ValidateAddress(address); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTContractMetadataURL, address, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetNFTContractMetadataResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) ReSyncMetadata(params ReSyncMetadataParams, address, tokenID string) (*ReSyncMetadataResponse, error) {
	if err := ValidateAddressAndTokenID(address, tokenID); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(ReSyncMetadataURL, address, tokenID, queryParams))
	if err != nil {
		return nil, err
	}

	res := &ReSyncMetadataResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) SyncNFTContract(params SyncNFTContractParams, address string) error {
	if err := ValidateAddress(address); err != nil {
		return err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return err
	}

	if _, _, err = c.doPutRequest(fmt.Sprintf(SyncNFTContractURL, address, queryParams)); err != nil {
		return err
	}

	return nil

}

func (c *client) GetNFTMetadata(params GetNFTMetadataParams, address, tokenID string) (*GetNFTMetadataResponse, error) {
	if err := ValidateAddressAndTokenID(address, tokenID); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTMetadataURL, address, tokenID, queryParams))
	if err != nil {
		return nil, err
	}

	res := &GetNFTMetadataResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTTokenIDOwners(params GetNFTTokenIDOwnersParams, address, tokenID string) (*GetNFTTokenIDOwnersResponse, error) {
	if err := ValidateAddressAndTokenID(address, tokenID); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTTokenIDOwnersURL, address, tokenID, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetNFTTokenIDOwnersResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client) GetNFTTransfers(params GetNFTTransfersParams, address, tokenID string) (*GetNFTTransfersResponse, error) {
	if err := ValidateAddressAndTokenID(address, tokenID); err != nil {
		return nil, err
	}
	queryParams, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doGetRequest(fmt.Sprintf(GetNFTTransfersURL, address, tokenID, queryParams))
	if err != nil {
		return nil, err
	}
	res := &GetNFTTransfersResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) doGetRequest(url string) ([]byte, int, error) {
	return c.doRequest(url, http.MethodGet)
}

func (c *client) doPutRequest(url string) ([]byte, int, error) {
	return c.doRequest(url, http.MethodPut)
}

func (c *client) doRequest(url, method string) ([]byte, int, error) {
	req, _ := http.NewRequest(method, fmt.Sprintf(NFTBaseURL, url), nil)

	if method == http.MethodGet {
		req.Header.Add("accept", "application/json")
	}
	req.Header.Add(APIKeyHeader, c.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, res.StatusCode, nil
}

func BuildQueryParams(params interface{}) (string, error) {
	paramsMap, err := StructToMap(params)
	if err != nil {
		return "", err
	}

	var (
		data          = url.Values{}
		addresses     = ""
		tokeAddresses = ""
	)

	for key, value := range paramsMap {
		switch key {
		case ChainField:
			data.Set(key, ValidateChain(value.(string)))
		case FormatField:
			data.Set(key, ValidateFormat(value.(string)))
		case DirectionField:
			data.Set(key, ValidateDirection(value.(string)))
		case MarketPlaceField:
			data.Set(key, ValidateMarketplace(value.(string)))
		case FilterField:
			data.Set(key, ValidateFilter(value.(string)))
		case FlagField:
			data.Set(key, ValidateFlag(value.(string)))
		case ModeField:
			data.Set(key, ValidateMode(value.(string)))
		case FromDateField, ToDateField:
			tm, err := time.Parse(time.RFC3339, value.(string))
			if err != nil {
				continue
			}
			if !tm.IsZero() {
				data.Set(key, tm.Format(time.RFC3339))
			}
		case AdressesField:
			switch val := value.(type) {
			case []interface{}:
				for _, v := range val {
					addresses += fmt.Sprintf("&%s=%s", key, v)
				}
			}
		case TokenAddressesField:
			switch val := value.(type) {
			case []interface{}:
				for _, v := range val {
					tokeAddresses += fmt.Sprintf("&%s=%s", key, v)
				}
			}
		default:
			switch v := value.(type) {
			case string:
				data.Set(key, v)
			case int:
				data.Set(key, strconv.Itoa(v))
			case float64:
				data.Set(key, strconv.Itoa(int(v)))
			}
		}
	}

	return data.Encode() + tokeAddresses + addresses, nil
}
