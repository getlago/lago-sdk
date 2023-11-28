# \FeesAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAllFees**](FeesAPI.md#FindAllFees) | **Get** /fees | List all fees
[**FindFee**](FeesAPI.md#FindFee) | **Get** /fees/{lago_id} | Retrieve a specific fee
[**UpdateFee**](FeesAPI.md#UpdateFee) | **Put** /fees/{lago_id} | Update a fee



## FindAllFees

> FeesPaginated FindAllFees(ctx).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Currency(currency).FeeType(feeType).BillableMetricCode(billableMetricCode).PaymentStatus(paymentStatus).CreatedAtFrom(createdAtFrom).CreatedAtTo(createdAtTo).SucceededAtFrom(succeededAtFrom).SucceededAtTo(succeededAtTo).FailedAtFrom(failedAtFrom).FailedAtTo(failedAtTo).RefundedAtFrom(refundedAtFrom).RefundedAtTo(refundedAtTo).Execute()

List all fees



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | Unique identifier assigned to the customer in your application. (optional)
    externalSubscriptionId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External subscription ID (optional)
    currency := openapiclient.Currency("AED") // Currency | Filter results by fee’s currency. (optional)
    feeType := "charge" // string | The fee type. Possible values are `add-on`, `charge`, `credit` or `subscription`. (optional)
    billableMetricCode := "bm_code" // string | Filter results by the `code` of the billable metric attached to the fee. Only applies to `charge` types. (optional)
    paymentStatus := "succeeded" // string | Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are `pending`, `succeeded`, `failed` and refunded`. (optional)
    createdAtFrom := time.Now() // time.Time | Filter results created after creation date and time in UTC. (optional)
    createdAtTo := time.Now() // time.Time | Filter results created before creation date and time in UTC. (optional)
    succeededAtFrom := time.Now() // time.Time | Filter results with payment success after creation date and time in UTC. (optional)
    succeededAtTo := time.Now() // time.Time | Filter results with payment success after creation date and time in UTC. (optional)
    failedAtFrom := time.Now() // time.Time | Filter results with payment failure after creation date and time in UTC. (optional)
    failedAtTo := time.Now() // time.Time | Filter results with payment failure after creation date and time in UTC. (optional)
    refundedAtFrom := time.Now() // time.Time | Filter results with payment refund after creation date and time in UTC. (optional)
    refundedAtTo := time.Now() // time.Time | Filter results with payment refund after creation date and time in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeesAPI.FindAllFees(context.Background()).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).ExternalSubscriptionId(externalSubscriptionId).Currency(currency).FeeType(feeType).BillableMetricCode(billableMetricCode).PaymentStatus(paymentStatus).CreatedAtFrom(createdAtFrom).CreatedAtTo(createdAtTo).SucceededAtFrom(succeededAtFrom).SucceededAtTo(succeededAtTo).FailedAtFrom(failedAtFrom).FailedAtTo(failedAtTo).RefundedAtFrom(refundedAtFrom).RefundedAtTo(refundedAtTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesAPI.FindAllFees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllFees`: FeesPaginated
    fmt.Fprintf(os.Stdout, "Response from `FeesAPI.FindAllFees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllFeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **externalCustomerId** | **string** | Unique identifier assigned to the customer in your application. | 
 **externalSubscriptionId** | **string** | External subscription ID | 
 **currency** | [**Currency**](Currency.md) | Filter results by fee’s currency. | 
 **feeType** | **string** | The fee type. Possible values are &#x60;add-on&#x60;, &#x60;charge&#x60;, &#x60;credit&#x60; or &#x60;subscription&#x60;. | 
 **billableMetricCode** | **string** | Filter results by the &#x60;code&#x60; of the billable metric attached to the fee. Only applies to &#x60;charge&#x60; types. | 
 **paymentStatus** | **string** | Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; and refunded&#x60;. | 
 **createdAtFrom** | **time.Time** | Filter results created after creation date and time in UTC. | 
 **createdAtTo** | **time.Time** | Filter results created before creation date and time in UTC. | 
 **succeededAtFrom** | **time.Time** | Filter results with payment success after creation date and time in UTC. | 
 **succeededAtTo** | **time.Time** | Filter results with payment success after creation date and time in UTC. | 
 **failedAtFrom** | **time.Time** | Filter results with payment failure after creation date and time in UTC. | 
 **failedAtTo** | **time.Time** | Filter results with payment failure after creation date and time in UTC. | 
 **refundedAtFrom** | **time.Time** | Filter results with payment refund after creation date and time in UTC. | 
 **refundedAtTo** | **time.Time** | Filter results with payment refund after creation date and time in UTC. | 

### Return type

[**FeesPaginated**](FeesPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindFee

> Fee FindFee(ctx, lagoId).Execute()

Retrieve a specific fee



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeesAPI.FindFee(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesAPI.FindFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindFee`: Fee
    fmt.Fprintf(os.Stdout, "Response from `FeesAPI.FindFee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Fee**](Fee.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFee

> Fee UpdateFee(ctx, lagoId).FeeUpdateInput(feeUpdateInput).Execute()

Update a fee



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.
    feeUpdateInput := *openapiclient.NewFeeUpdateInput(*openapiclient.NewFeeUpdateInputFee("succeeded")) // FeeUpdateInput | Fee payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeesAPI.UpdateFee(context.Background(), lagoId).FeeUpdateInput(feeUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeesAPI.UpdateFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFee`: Fee
    fmt.Fprintf(os.Stdout, "Response from `FeesAPI.UpdateFee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feeUpdateInput** | [**FeeUpdateInput**](FeeUpdateInput.md) | Fee payload | 

### Return type

[**Fee**](Fee.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

