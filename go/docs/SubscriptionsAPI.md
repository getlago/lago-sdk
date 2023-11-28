# \SubscriptionsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsAPI.md#CreateSubscription) | **Post** /subscriptions | Assign a plan to a customer
[**DestroySubscription**](SubscriptionsAPI.md#DestroySubscription) | **Delete** /subscriptions/{external_id} | Terminate a subscription
[**FindAllSubscriptions**](SubscriptionsAPI.md#FindAllSubscriptions) | **Get** /subscriptions | List all subscriptions
[**FindSubscription**](SubscriptionsAPI.md#FindSubscription) | **Get** /subscriptions/{external_id} | Retrieve a subscription
[**UpdateSubscription**](SubscriptionsAPI.md#UpdateSubscription) | **Put** /subscriptions/{external_id} | Update a subscription



## CreateSubscription

> Subscription CreateSubscription(ctx).SubscriptionCreateInput(subscriptionCreateInput).Execute()

Assign a plan to a customer



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
    subscriptionCreateInput := *openapiclient.NewSubscriptionCreateInput(*openapiclient.NewSubscriptionCreateInputSubscription("5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba", "premium", "my_sub_1234567890")) // SubscriptionCreateInput | Subscription payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.CreateSubscription(context.Background()).SubscriptionCreateInput(subscriptionCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreateInput** | [**SubscriptionCreateInput**](SubscriptionCreateInput.md) | Subscription payload | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySubscription

> Subscription DestroySubscription(ctx, externalId).Status(status).Execute()

Terminate a subscription



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
    externalId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing subscription
    status := "pending" // string | If the field is not defined, Lago will terminate only `active` subscriptions. However, if you wish to cancel a `pending` subscription, please ensure that you include `status=pending` in your request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.DestroySubscription(context.Background(), externalId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.DestroySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroySubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.DestroySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | External ID of the existing subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | If the field is not defined, Lago will terminate only &#x60;active&#x60; subscriptions. However, if you wish to cancel a &#x60;pending&#x60; subscription, please ensure that you include &#x60;status&#x3D;pending&#x60; in your request. | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllSubscriptions

> SubscriptionsPaginated FindAllSubscriptions(ctx).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).PlanCode(planCode).Status(status).Execute()

List all subscriptions



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
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | The customer external unique identifier (provided by your own application) (optional)
    planCode := "premium" // string | The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans. (optional)
    status := []string{"Status_example"} // []string | If the field is not defined, Lago will return only `active` subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: `pending`, `canceled`, `terminated`, `active`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.FindAllSubscriptions(context.Background()).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).PlanCode(planCode).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.FindAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllSubscriptions`: SubscriptionsPaginated
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.FindAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **externalCustomerId** | **string** | The customer external unique identifier (provided by your own application) | 
 **planCode** | **string** | The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans. | 
 **status** | **[]string** | If the field is not defined, Lago will return only &#x60;active&#x60; subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: &#x60;pending&#x60;, &#x60;canceled&#x60;, &#x60;terminated&#x60;, &#x60;active&#x60;. | 

### Return type

[**SubscriptionsPaginated**](SubscriptionsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSubscription

> Subscription FindSubscription(ctx, externalId).Execute()

Retrieve a subscription



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
    externalId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.FindSubscription(context.Background(), externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.FindSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.FindSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | External ID of the existing subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> Subscription UpdateSubscription(ctx, externalId).SubscriptionUpdateInput(subscriptionUpdateInput).Execute()

Update a subscription



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
    externalId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | External ID of the existing subscription
    subscriptionUpdateInput := *openapiclient.NewSubscriptionUpdateInput(*openapiclient.NewSubscriptionUpdateInputSubscription()) // SubscriptionUpdateInput | Update an existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.UpdateSubscription(context.Background(), externalId).SubscriptionUpdateInput(subscriptionUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | External ID of the existing subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionUpdateInput** | [**SubscriptionUpdateInput**](SubscriptionUpdateInput.md) | Update an existing subscription | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

