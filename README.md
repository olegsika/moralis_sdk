# Moralis SDK

### Installation 
    go get github.com/olegsika/moralis_sdk

### Api Documentation

https://docs.moralis.io/reference/nft-api

### Example
```go
client := moralis.NewClient("your api key")

res, err := client.GetWalletNFTs(params, address)

fmt.Println(err) // error
fmt.Println(res) // GetWalletNFTsResponse
```

### Hint

All the structs you can see on nft_structs.go

All the methods you can see on interfaces.go

*This is not an official Moralis product.*
