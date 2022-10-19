package moralissdk

import (
	"errors"
	"testing"
)

func TestValidateBlockNumberOrHash(t *testing.T) {
	tCases := []struct {
		name              string
		blockNumberOrHash string
		expectedErr       error
	}{
		{
			name:              "should be success",
			blockNumberOrHash: "test",
		},
		{
			name:              "should be error",
			blockNumberOrHash: "",
			expectedErr:       errors.New("blockNumberOrHash is required"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := ValidateBlockNumberOrHash(tc.blockNumberOrHash)
			if !IsErrorEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
			}
		})
	}
}
func TestValidateAddress(t *testing.T) {
	tCases := []struct {
		name        string
		address     string
		expectedErr error
	}{
		{
			name:    "should be success",
			address: "test",
		},
		{
			name:        "should be error",
			address:     "",
			expectedErr: errors.New("address is required"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := ValidateAddress(tc.address)
			if !IsErrorEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
			}
		})
	}
}

func TestValidateAddressAndTokenID(t *testing.T) {
	tCases := []struct {
		name        string
		address     string
		tokenID     string
		expectedErr error
	}{
		{
			name:    "should be success",
			address: "test",
			tokenID: "test",
		},
		{
			name:        "should be error. Empty Address",
			address:     "",
			tokenID:     "test",
			expectedErr: errors.New("address is required"),
		},
		{
			name:        "should be error. Empty Token ID",
			address:     "test",
			tokenID:     "",
			expectedErr: errors.New("tokenID is required"),
		},
		{
			name:        "should be error. Empty Address and Token ID",
			address:     "",
			tokenID:     "",
			expectedErr: errors.New("address is required"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := ValidateAddressAndTokenID(tc.address, tc.tokenID)
			if !IsErrorEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
			}
		})
	}
}

func TestValidateChain(t *testing.T) {
	tCases := map[string]string{
		"eth":               "eth",
		"0x1":               "0x1",
		"goerli":            "goerli",
		"0x5":               "0x5",
		"sepolia":           "sepolia",
		"0xaa36a7":          "0xaa36a7",
		"polygon":           "polygon",
		"0x89":              "0x89",
		"mumbai":            "mumbai",
		"0x13881":           "0x13881",
		"bsc":               "bsc",
		"0x38":              "0x38",
		"bsc testnet":       "bsc testnet",
		"0x61":              "0x61",
		"avalanche":         "avalanche",
		"0xa86a":            "0xa86a",
		"avalanche testnet": "avalanche testnet",
		"0xa869":            "0xa869",
		"fantom":            "fantom",
		"0xfa":              "0xfa",
		"cronos":            "cronos",
		"0x19":              "0x19",
		"cronos testnet":    "cronos testnet",
		"0x152":             "0x152",
		"":                  "eth",
		"test":              "eth",
	}

	for key, expectedRes := range tCases {
		res := ValidateChain(key)

		if res != expectedRes {
			t.Errorf("expected chain %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateFormat(t *testing.T) {
	tCases := map[string]string{
		"decimal": "decimal",
		"hex":     "hex",
		"":        "decimal",
		"test":    "decimal",
	}

	for key, expectedRes := range tCases {
		res := ValidateFormat(key)

		if res != expectedRes {
			t.Errorf("expected format %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateDirection(t *testing.T) {
	tCases := map[string]string{
		"both": "both",
		"to":   "to",
		"from": "from",
		"":     "both",
		"test": "both",
	}

	for key, expectedRes := range tCases {
		res := ValidateDirection(key)

		if res != expectedRes {
			t.Errorf("expected direction %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateMarketplace(t *testing.T) {
	tCases := map[string]string{
		"opensea": "opensea",
		"":        "opensea",
		"test":    "opensea",
	}

	for key, expectedRes := range tCases {
		res := ValidateMarketplace(key)

		if res != expectedRes {
			t.Errorf("expected marketplace %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateFilter(t *testing.T) {
	tCases := map[string]string{
		"name":                          "name",
		"description":                   "description",
		"attributes":                    "attributes",
		"global":                        "global",
		"name, description":             "name, description",
		"name, attributes":              "name, attributes",
		"description, attributes":       "description, attributes",
		"name, description, attributes": "name, description, attributes",
		"":                              "global",
		"test":                          "global",
	}

	for key, expectedRes := range tCases {
		res := ValidateFilter(key)

		if res != expectedRes {
			t.Errorf("expected filter %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateFlag(t *testing.T) {
	tCases := map[string]string{
		"uri":      "uri",
		"metadata": "metadata",
		"":         "uri",
		"test":     "uri",
	}

	for key, expectedRes := range tCases {
		res := ValidateFlag(key)

		if res != expectedRes {
			t.Errorf("expected flag %s but actual %s", expectedRes, res)
		}
	}
}

func TestValidateMode(t *testing.T) {
	tCases := map[string]string{
		"async": "async",
		"sync":  "sync",
		"":      "async",
		"test":  "async",
	}

	for key, expectedRes := range tCases {
		res := ValidateMode(key)

		if res != expectedRes {
			t.Errorf("expected flag %s but actual %s", expectedRes, res)
		}
	}
}
