# \CustomersAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersAPI.md#CreateCustomer) | **Post** /customers | Create a customer
[**DeleteAppliedCoupon**](CustomersAPI.md#DeleteAppliedCoupon) | **Delete** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
[**DestroyCustomer**](CustomersAPI.md#DestroyCustomer) | **Delete** /customers/{external_id} | Delete a customer
[**FindAllCustomerPastUsage**](CustomersAPI.md#FindAllCustomerPastUsage) | **Get** /customers/{external_customer_id}/past_usage | Retrieve customer past usage
[**FindAllCustomers**](CustomersAPI.md#FindAllCustomers) | **Get** /customers | List all customers
[**FindCustomer**](CustomersAPI.md#FindCustomer) | **Get** /customers/{external_id} | Retrieve a customer
[**FindCustomerCurrentUsage**](CustomersAPI.md#FindCustomerCurrentUsage) | **Get** /customers/{external_customer_id}/current_usage | Retrieve customer current usage
[**GenerateCustomerCheckoutURL**](CustomersAPI.md#GenerateCustomerCheckoutURL) | **Post** /customers/{external_customer_id}/checkout_url | Generate a Customer Payment Provider Checkout URL
[**GetCustomerPortalUrl**](CustomersAPI.md#GetCustomerPortalUrl) | **Get** /customers/{external_customer_id}/portal_url | Get a customer portal URL



## CreateCustomer

> Customer CreateCustomer(ctx).CustomerCreateInput(customerCreateInput).Execute()

Create a customer



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
    customerCreateInput := *openapiclient.NewCustomerCreateInput(*openapiclient.NewCustomerCreateInputCustomer("5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba")) // CustomerCreateInput | Customer payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).CustomerCreateInput(customerCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCreateInput** | [**CustomerCreateInput**](CustomerCreateInput.md) | Customer payload | 

### Return type

[**Customer**](Customer.md)

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
    resp, r, err := apiClient.CustomersAPI.DeleteAppliedCoupon(context.Background(), externalCustomerId, appliedCouponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.DeleteAppliedCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppliedCoupon`: AppliedCoupon
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.DeleteAppliedCoupon`: %v\n", resp)
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


## DestroyCustomer

> Customer DestroyCustomer(ctx, externalId).Execute()

Delete a customer



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
    externalId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing customer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.DestroyCustomer(context.Background(), externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.DestroyCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.DestroyCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | External ID of the existing customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllCustomerPastUsage

> CustomerPastUsage FindAllCustomerPastUsage(ctx, externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Page(page).PerPage(perPage).BillableMetricCode(billableMetricCode).PeriodsCount(periodsCount).Execute()

Retrieve customer past usage



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
    externalSubscriptionId := "sub_1234567890" // string | The unique identifier of the subscription within your application.
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)
    billableMetricCode := "cpu" // string | Billable metric code filter to apply to the charge usage (optional)
    periodsCount := int32(5) // int32 | Number of past billing period to returns in the result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.FindAllCustomerPastUsage(context.Background(), externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Page(page).PerPage(perPage).BillableMetricCode(billableMetricCode).PeriodsCount(periodsCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.FindAllCustomerPastUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllCustomerPastUsage`: CustomerPastUsage
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.FindAllCustomerPastUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCustomerId** | **string** | The customer external unique identifier (provided by your own application). | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllCustomerPastUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalSubscriptionId** | **string** | The unique identifier of the subscription within your application. | 
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **billableMetricCode** | **string** | Billable metric code filter to apply to the charge usage | 
 **periodsCount** | **int32** | Number of past billing period to returns in the result | 

### Return type

[**CustomerPastUsage**](CustomerPastUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllCustomers

> CustomersPaginated FindAllCustomers(ctx).Page(page).PerPage(perPage).Execute()

List all customers



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
    resp, r, err := apiClient.CustomersAPI.FindAllCustomers(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.FindAllCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllCustomers`: CustomersPaginated
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.FindAllCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**CustomersPaginated**](CustomersPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCustomer

> Customer FindCustomer(ctx, externalId).Execute()

Retrieve a customer



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
    externalId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing customer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.FindCustomer(context.Background(), externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.FindCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.FindCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | External ID of the existing customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCustomerCurrentUsage

> CustomerUsage FindCustomerCurrentUsage(ctx, externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Execute()

Retrieve customer current usage



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
    externalSubscriptionId := "sub_1234567890" // string | The unique identifier of the subscription within your application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.FindCustomerCurrentUsage(context.Background(), externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.FindCustomerCurrentUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCustomerCurrentUsage`: CustomerUsage
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.FindCustomerCurrentUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCustomerId** | **string** | The customer external unique identifier (provided by your own application). | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCustomerCurrentUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalSubscriptionId** | **string** | The unique identifier of the subscription within your application. | 

### Return type

[**CustomerUsage**](CustomerUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCustomerCheckoutURL

> GenerateCustomerCheckoutURL200Response GenerateCustomerCheckoutURL(ctx, externalCustomerId).Execute()

Generate a Customer Payment Provider Checkout URL



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.GenerateCustomerCheckoutURL(context.Background(), externalCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GenerateCustomerCheckoutURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCustomerCheckoutURL`: GenerateCustomerCheckoutURL200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GenerateCustomerCheckoutURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCustomerId** | **string** | The customer external unique identifier (provided by your own application). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCustomerCheckoutURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerateCustomerCheckoutURL200Response**](GenerateCustomerCheckoutURL200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerPortalUrl

> GetCustomerPortalUrl200Response GetCustomerPortalUrl(ctx, externalCustomerId).Execute()

Get a customer portal URL



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
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing customer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.GetCustomerPortalUrl(context.Background(), externalCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerPortalUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerPortalUrl`: GetCustomerPortalUrl200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerPortalUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCustomerId** | **string** | External ID of the existing customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerPortalUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomerPortalUrl200Response**](GetCustomerPortalUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

