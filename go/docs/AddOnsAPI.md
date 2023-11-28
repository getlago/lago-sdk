# \AddOnsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddOn**](AddOnsAPI.md#CreateAddOn) | **Post** /add_ons | Create an add-on
[**DestroyAddOn**](AddOnsAPI.md#DestroyAddOn) | **Delete** /add_ons/{code} | Delete an add-on
[**FindAddOn**](AddOnsAPI.md#FindAddOn) | **Get** /add_ons/{code} | Retrieve an add-on
[**FindAllAddOns**](AddOnsAPI.md#FindAllAddOns) | **Get** /add_ons | List all add-ons
[**UpdateAddOn**](AddOnsAPI.md#UpdateAddOn) | **Put** /add_ons/{code} | Update an add-on



## CreateAddOn

> AddOn CreateAddOn(ctx).AddOnCreateInput(addOnCreateInput).Execute()

Create an add-on



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
    addOnCreateInput := *openapiclient.NewAddOnCreateInput(*openapiclient.NewAddOnCreateInputAddOn("Setup Fee", "setup_fee", int32(50000), openapiclient.Currency("AED"))) // AddOnCreateInput | Add-on payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddOnsAPI.CreateAddOn(context.Background()).AddOnCreateInput(addOnCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddOnsAPI.CreateAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAddOn`: AddOn
    fmt.Fprintf(os.Stdout, "Response from `AddOnsAPI.CreateAddOn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addOnCreateInput** | [**AddOnCreateInput**](AddOnCreateInput.md) | Add-on payload | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyAddOn

> AddOn DestroyAddOn(ctx, code).Execute()

Delete an add-on



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
    code := "setup_fee" // string | Unique code used to identify the add-on.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddOnsAPI.DestroyAddOn(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddOnsAPI.DestroyAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyAddOn`: AddOn
    fmt.Fprintf(os.Stdout, "Response from `AddOnsAPI.DestroyAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the add-on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAddOn

> AddOn FindAddOn(ctx, code).Execute()

Retrieve an add-on



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
    code := "setup_fee" // string | Unique code used to identify the add-on.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddOnsAPI.FindAddOn(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddOnsAPI.FindAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAddOn`: AddOn
    fmt.Fprintf(os.Stdout, "Response from `AddOnsAPI.FindAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the add-on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllAddOns

> AddOnsPaginated FindAllAddOns(ctx).Page(page).PerPage(perPage).Execute()

List all add-ons



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
    resp, r, err := apiClient.AddOnsAPI.FindAllAddOns(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddOnsAPI.FindAllAddOns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllAddOns`: AddOnsPaginated
    fmt.Fprintf(os.Stdout, "Response from `AddOnsAPI.FindAllAddOns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllAddOnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**AddOnsPaginated**](AddOnsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddOn

> AddOn UpdateAddOn(ctx, code).AddOnUpdateInput(addOnUpdateInput).Execute()

Update an add-on



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
    code := "setup_fee" // string | Unique code used to identify the add-on.
    addOnUpdateInput := *openapiclient.NewAddOnUpdateInput(*openapiclient.NewAddOnBaseInput()) // AddOnUpdateInput | Add-on payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddOnsAPI.UpdateAddOn(context.Background(), code).AddOnUpdateInput(addOnUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddOnsAPI.UpdateAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddOn`: AddOn
    fmt.Fprintf(os.Stdout, "Response from `AddOnsAPI.UpdateAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the add-on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addOnUpdateInput** | [**AddOnUpdateInput**](AddOnUpdateInput.md) | Add-on payload | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

