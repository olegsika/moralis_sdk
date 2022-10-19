package moralissdk

import "errors"

func ValidateBlockNumberOrHash(blockNumberOrHash string) error {
	if blockNumberOrHash == "" {
		return errors.New("blockNumberOrHash is required")
	}

	return nil
}

func ValidateAddress(address string) error {
	if address == "" {
		return errors.New("address is required")
	}

	return nil
}

func ValidateAddressAndTokenID(address, tokenID string) error {
	if address == "" {
		return errors.New("address is required")
	}
	if tokenID == "" {
		return errors.New("tokenID is required")
	}

	return nil
}

func ValidateChain(chain string) string {
	availableChains := []string{
		"eth",
		"0x1",
		"goerli",
		"0x5",
		"sepolia",
		"0xaa36a7",
		"polygon",
		"0x89",
		"mumbai",
		"0x13881",
		"bsc",
		"0x38",
		"bsc testnet",
		"0x61",
		"avalanche",
		"0xa86a",
		"avalanche testnet",
		"0xa869",
		"fantom",
		"0xfa",
		"cronos",
		"0x19",
		"cronos testnet",
		"0x152",
	}

	if !InListString(chain, availableChains) {
		return "eth"
	}

	return chain
}

func ValidateFormat(format string) string {
	availableFormats := []string{"decimal", "hex"}

	if !InListString(format, availableFormats) {
		return "decimal"
	}

	return format
}

func ValidateDirection(direction string) string {
	availableDirections := []string{"both", "to", "from"}

	if !InListString(direction, availableDirections) {
		return "both"
	}

	return direction
}

func ValidateMarketplace(marketplace string) string {
	availableMarketplaces := []string{"opensea"}

	if !InListString(marketplace, availableMarketplaces) {
		return "opensea"
	}

	return marketplace
}

func ValidateFilter(filter string) string {
	availableFilters := []string{
		"name",
		"description",
		"attributes",
		"global",
		"name, description",
		"name, attributes",
		"description, attributes",
		"name, description, attributes",
	}

	if !InListString(filter, availableFilters) {
		return "global"
	}

	return filter
}

func ValidateFlag(flag string) string {
	availableFlags := []string{"uri", "metadata"}

	if !InListString(flag, availableFlags) {
		return "uri"
	}

	return flag
}

func ValidateMode(mode string) string {
	availableModes := []string{"async", "sync"}

	if !InListString(mode, availableModes) {
		return "async"
	}

	return mode
}
