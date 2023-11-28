# \WalletsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWallet**](WalletsAPI.md#CreateWallet) | **Post** /wallets | Create a wallet
[**CreateWalletTransaction**](WalletsAPI.md#CreateWalletTransaction) | **Post** /wallet_transactions | Top up a wallet
[**DestroyWallet**](WalletsAPI.md#DestroyWallet) | **Delete** /wallets/{lago_id} | Terminate a wallet
[**FindAllWalletTransactions**](WalletsAPI.md#FindAllWalletTransactions) | **Get** /wallets/{lago_id}/wallet_transactions | List all wallet transactions
[**FindAllWallets**](WalletsAPI.md#FindAllWallets) | **Get** /wallets | List all wallets
[**FindWallet**](WalletsAPI.md#FindWallet) | **Get** /wallets/{lago_id} | Retrieve a wallet
[**UpdateWallet**](WalletsAPI.md#UpdateWallet) | **Put** /wallets/{lago_id} | Update a wallet



## CreateWallet

> Wallet CreateWallet(ctx).WalletCreateInput(walletCreateInput).Execute()

Create a wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    walletCreateInput := *openapiclient.NewWalletCreateInput() // WalletCreateInput | Wallet payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.CreateWallet(context.Background()).WalletCreateInput(walletCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWallet`: Wallet
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletCreateInput** | [**WalletCreateInput**](WalletCreateInput.md) | Wallet payload | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWalletTransaction

> WalletTransactions CreateWalletTransaction(ctx).WalletTransactionCreateInput(walletTransactionCreateInput).Execute()

Top up a wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    walletTransactionCreateInput := *openapiclient.NewWalletTransactionCreateInput(*openapiclient.NewWalletTransactionCreateInputWalletTransaction("1a901a90-1a90-1a90-1a90-1a901a901a90")) // WalletTransactionCreateInput | Wallet transaction payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.CreateWalletTransaction(context.Background()).WalletTransactionCreateInput(walletTransactionCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWalletTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWalletTransaction`: WalletTransactions
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWalletTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletTransactionCreateInput** | [**WalletTransactionCreateInput**](WalletTransactionCreateInput.md) | Wallet transaction payload | 

### Return type

[**WalletTransactions**](WalletTransactions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWallet

> Wallet DestroyWallet(ctx, lagoId).Execute()

Terminate a wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.DestroyWallet(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DestroyWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyWallet`: Wallet
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.DestroyWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllWalletTransactions

> WalletTransactionsPaginated FindAllWalletTransactions(ctx, lagoId).Page(page).PerPage(perPage).Status(status).TransactionType(transactionType).Execute()

List all wallet transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)
    status := "pending" // string | The status of the wallet transaction. Possible values are `pending` or `settled`. (optional)
    transactionType := "inbound" // string | The transaction type of the wallet transaction. Possible values are `inbound` (increasing the wallet balance) or `outbound` (decreasing the wallet balance). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.FindAllWalletTransactions(context.Background(), lagoId).Page(page).PerPage(perPage).Status(status).TransactionType(transactionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.FindAllWalletTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllWalletTransactions`: WalletTransactionsPaginated
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.FindAllWalletTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **status** | **string** | The status of the wallet transaction. Possible values are &#x60;pending&#x60; or &#x60;settled&#x60;. | 
 **transactionType** | **string** | The transaction type of the wallet transaction. Possible values are &#x60;inbound&#x60; (increasing the wallet balance) or &#x60;outbound&#x60; (decreasing the wallet balance). | 

### Return type

[**WalletTransactionsPaginated**](WalletTransactionsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllWallets

> WalletsPaginated FindAllWallets(ctx).ExternalCustomerId(externalCustomerId).Page(page).PerPage(perPage).Execute()

List all wallets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | The customer external unique identifier (provided by your own application).
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.FindAllWallets(context.Background()).ExternalCustomerId(externalCustomerId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.FindAllWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllWallets`: WalletsPaginated
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.FindAllWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalCustomerId** | **string** | The customer external unique identifier (provided by your own application). | 
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**WalletsPaginated**](WalletsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWallet

> Wallet FindWallet(ctx, lagoId).Execute()

Retrieve a wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.FindWallet(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.FindWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWallet`: Wallet
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.FindWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWallet

> Wallet UpdateWallet(ctx, lagoId).WalletUpdateInput(walletUpdateInput).Execute()

Update a wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.
    walletUpdateInput := *openapiclient.NewWalletUpdateInput(*openapiclient.NewWalletUpdateInputWallet()) // WalletUpdateInput | Wallet update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsAPI.UpdateWallet(context.Background(), lagoId).WalletUpdateInput(walletUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWallet`: Wallet
    fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletUpdateInput** | [**WalletUpdateInput**](WalletUpdateInput.md) | Wallet update payload | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

