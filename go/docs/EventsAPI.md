# \EventsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchEvents**](EventsAPI.md#CreateBatchEvents) | **Post** /events/batch | Batch multiple events
[**CreateEvent**](EventsAPI.md#CreateEvent) | **Post** /events | Send usage events
[**EventEstimateFees**](EventsAPI.md#EventEstimateFees) | **Post** /events/estimate_fees | Estimate fees for a pay in advance charge
[**FindEvent**](EventsAPI.md#FindEvent) | **Get** /events/{transaction_id} | Retrieve a specific event



## CreateBatchEvents

> CreateBatchEvents(ctx).EventBatchInput(eventBatchInput).Execute()

Batch multiple events



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
    eventBatchInput := *openapiclient.NewEventBatchInput(*openapiclient.NewEventBatchInputEvent("transaction_1234567890", []string{"ExternalSubscriptionIds_example"}, "storage")) // EventBatchInput | Batch events payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.CreateBatchEvents(context.Background()).EventBatchInput(eventBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateBatchEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventBatchInput** | [**EventBatchInput**](EventBatchInput.md) | Batch events payload | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvent

> Event CreateEvent(ctx).EventInput(eventInput).Execute()

Send usage events



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
    eventInput := *openapiclient.NewEventInput(*openapiclient.NewEventInputEvent("transaction_1234567890", "storage")) // EventInput | Event payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.CreateEvent(context.Background()).EventInput(eventInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvent`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.CreateEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventInput** | [**EventInput**](EventInput.md) | Event payload | 

### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventEstimateFees

> Fees EventEstimateFees(ctx).EventEstimateFeesInput(eventEstimateFeesInput).Execute()

Estimate fees for a pay in advance charge



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
    eventEstimateFeesInput := *openapiclient.NewEventEstimateFeesInput(*openapiclient.NewEventEstimateFeesInputEvent("storage")) // EventEstimateFeesInput | Event estimate payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventEstimateFees(context.Background()).EventEstimateFeesInput(eventEstimateFeesInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventEstimateFees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventEstimateFees`: Fees
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventEstimateFees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventEstimateFeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventEstimateFeesInput** | [**EventEstimateFeesInput**](EventEstimateFeesInput.md) | Event estimate payload | 

### Return type

[**Fees**](Fees.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindEvent

> Event FindEvent(ctx, transactionId).Execute()

Retrieve a specific event



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
    transactionId := "transaction_1234567890" // string | This field represents the unique identifier sent for this specific event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.FindEvent(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.FindEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindEvent`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.FindEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | This field represents the unique identifier sent for this specific event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

