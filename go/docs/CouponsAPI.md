# \CouponsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCoupon**](CouponsAPI.md#ApplyCoupon) | **Post** /applied_coupons | Apply a coupon to a customer
[**CreateCoupon**](CouponsAPI.md#CreateCoupon) | **Post** /coupons | Create a coupon
[**DeleteAppliedCoupon**](CouponsAPI.md#DeleteAppliedCoupon) | **Delete** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
[**DestroyCoupon**](CouponsAPI.md#DestroyCoupon) | **Delete** /coupons/{code} | Delete a coupon
[**FindAllAppliedCoupons**](CouponsAPI.md#FindAllAppliedCoupons) | **Get** /applied_coupons | List all applied coupons
[**FindAllCoupons**](CouponsAPI.md#FindAllCoupons) | **Get** /coupons | List all coupons
[**FindCoupon**](CouponsAPI.md#FindCoupon) | **Get** /coupons/{code} | Retrieve a coupon
[**UpdateCoupon**](CouponsAPI.md#UpdateCoupon) | **Put** /coupons/{code} | Update a coupon



## ApplyCoupon

> AppliedCoupon ApplyCoupon(ctx).AppliedCouponInput(appliedCouponInput).Execute()

Apply a coupon to a customer



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
    appliedCouponInput := *openapiclient.NewAppliedCouponInput(*openapiclient.NewAppliedCouponInputAppliedCoupon("5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba", "startup_deal")) // AppliedCouponInput | Apply coupon payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.ApplyCoupon(context.Background()).AppliedCouponInput(appliedCouponInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.ApplyCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyCoupon`: AppliedCoupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.ApplyCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedCouponInput** | [**AppliedCouponInput**](AppliedCouponInput.md) | Apply coupon payload | 

### Return type

[**AppliedCoupon**](AppliedCoupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupon

> Coupon CreateCoupon(ctx).CouponCreateInput(couponCreateInput).Execute()

Create a coupon



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
    couponCreateInput := *openapiclient.NewCouponCreateInput(*openapiclient.NewCouponCreateInputCoupon("Startup Deal", "startup_deal", "fixed_amount", "recurring")) // CouponCreateInput | Coupon payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.CreateCoupon(context.Background()).CouponCreateInput(couponCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CreateCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCoupon`: Coupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CreateCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCreateInput** | [**CouponCreateInput**](CouponCreateInput.md) | Coupon payload | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppliedCoupon

> AppliedCoupon DeleteAppliedCoupon(ctx, externalCustomerId, appliedCouponId).Execute()

Delete an applied coupon



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
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | The customer external unique identifier (provided by your own application)
    appliedCouponId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier of the applied coupon, created by Lago.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.DeleteAppliedCoupon(context.Background(), externalCustomerId, appliedCouponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.DeleteAppliedCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppliedCoupon`: AppliedCoupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.DeleteAppliedCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
**appliedCouponId** | **string** | Unique identifier of the applied coupon, created by Lago. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppliedCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppliedCoupon**](AppliedCoupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCoupon

> Coupon DestroyCoupon(ctx, code).Execute()

Delete a coupon



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
    code := "startup_deal" // string | Unique code used to identify the coupon.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.DestroyCoupon(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.DestroyCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyCoupon`: Coupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.DestroyCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllAppliedCoupons

> AppliedCouponsPaginated FindAllAppliedCoupons(ctx).Page(page).PerPage(perPage).Status(status).ExternalCustomerId(externalCustomerId).Execute()

List all applied coupons



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
    status := "active" // string | The status of the coupon. Can be either `active` or `terminated`. (optional)
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | The customer external unique identifier (provided by your own application) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.FindAllAppliedCoupons(context.Background()).Page(page).PerPage(perPage).Status(status).ExternalCustomerId(externalCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.FindAllAppliedCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllAppliedCoupons`: AppliedCouponsPaginated
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.FindAllAppliedCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllAppliedCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **status** | **string** | The status of the coupon. Can be either &#x60;active&#x60; or &#x60;terminated&#x60;. | 
 **externalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 

### Return type

[**AppliedCouponsPaginated**](AppliedCouponsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllCoupons

> CouponsPaginated FindAllCoupons(ctx).Page(page).PerPage(perPage).Execute()

List all coupons



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
    resp, r, err := apiClient.CouponsAPI.FindAllCoupons(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.FindAllCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllCoupons`: CouponsPaginated
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.FindAllCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**CouponsPaginated**](CouponsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCoupon

> Coupon FindCoupon(ctx, code).Execute()

Retrieve a coupon



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
    code := "startup_deal" // string | Unique code used to identify the coupon.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.FindCoupon(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.FindCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCoupon`: Coupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.FindCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> Coupon UpdateCoupon(ctx, code).CouponUpdateInput(couponUpdateInput).Execute()

Update a coupon



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
    code := "startup_deal" // string | Unique code used to identify the coupon.
    couponUpdateInput := *openapiclient.NewCouponUpdateInput(*openapiclient.NewCouponBaseInput()) // CouponUpdateInput | Coupon payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.UpdateCoupon(context.Background(), code).CouponUpdateInput(couponUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.UpdateCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCoupon`: Coupon
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.UpdateCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Unique code used to identify the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **couponUpdateInput** | [**CouponUpdateInput**](CouponUpdateInput.md) | Coupon payload | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

