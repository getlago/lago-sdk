# \TaxesAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTax**](TaxesAPI.md#CreateTax) | **Post** /taxes | Create a tax
[**DestroyTax**](TaxesAPI.md#DestroyTax) | **Delete** /taxes/{code} | Delete a tax
[**FindAllTaxes**](TaxesAPI.md#FindAllTaxes) | **Get** /taxes | List all taxes
[**FindTax**](TaxesAPI.md#FindTax) | **Get** /taxes/{code} | Retrieve a Tax
[**UpdateTax**](TaxesAPI.md#UpdateTax) | **Put** /taxes/{code} | Update a tax



## CreateTax

> Tax CreateTax(ctx).TaxCreateInput(taxCreateInput).Execute()

Create a tax



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
    taxCreateInput := *openapiclient.NewTaxCreateInput(*openapiclient.NewTaxCreateInputTax("TVA", "french_standard_vat", "20.0")) // TaxCreateInput | Tax creation payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesAPI.CreateTax(context.Background()).TaxCreateInput(taxCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesAPI.CreateTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTax`: Tax
    fmt.Fprintf(os.Stdout, "Response from `TaxesAPI.CreateTax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxCreateInput** | [**TaxCreateInput**](TaxCreateInput.md) | Tax creation payload | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTax

> Tax DestroyTax(ctx, code).Execute()

Delete a tax



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
    code := "french_standard_vat" // string | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesAPI.DestroyTax(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesAPI.DestroyTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyTax`: Tax
    fmt.Fprintf(os.Stdout, "Response from `TaxesAPI.DestroyTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllTaxes

> TaxesPaginated FindAllTaxes(ctx).Page(page).PerPage(perPage).Execute()

List all taxes



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
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesAPI.FindAllTaxes(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesAPI.FindAllTaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllTaxes`: TaxesPaginated
    fmt.Fprintf(os.Stdout, "Response from `TaxesAPI.FindAllTaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**TaxesPaginated**](TaxesPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTax

> Tax FindTax(ctx, code).Execute()

Retrieve a Tax



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
    code := "french_standard_vat" // string | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesAPI.FindTax(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesAPI.FindTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindTax`: Tax
    fmt.Fprintf(os.Stdout, "Response from `TaxesAPI.FindTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTax

> Tax UpdateTax(ctx, code).TaxUpdateInput(taxUpdateInput).Execute()

Update a tax



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
    code := "french_standard_vat" // string | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.
    taxUpdateInput := *openapiclient.NewTaxUpdateInput(*openapiclient.NewTaxBaseInput()) // TaxUpdateInput | Tax update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesAPI.UpdateTax(context.Background(), code).TaxUpdateInput(taxUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesAPI.UpdateTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTax`: Tax
    fmt.Fprintf(os.Stdout, "Response from `TaxesAPI.UpdateTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxUpdateInput** | [**TaxUpdateInput**](TaxUpdateInput.md) | Tax update payload | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

