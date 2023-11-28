# \WebhookEndpointsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookEndpoint**](WebhookEndpointsAPI.md#CreateWebhookEndpoint) | **Post** /webhook_endpoints | Create a webhook_endpoint
[**DestroyWebhookEndpoint**](WebhookEndpointsAPI.md#DestroyWebhookEndpoint) | **Delete** /webhook_endpoints/{lago_id} | Delete a webhook endpoint
[**FindAllWebhookEndpoints**](WebhookEndpointsAPI.md#FindAllWebhookEndpoints) | **Get** /webhook_endpoints | List all webhook endpoints
[**FindWebhookEndpoint**](WebhookEndpointsAPI.md#FindWebhookEndpoint) | **Get** /webhook_endpoints/{lago_id} | Retrieve a webhook endpoint
[**UpdateWebhookEndpoint**](WebhookEndpointsAPI.md#UpdateWebhookEndpoint) | **Put** /webhook_endpoints/{lago_id} | Update a webhook endpoint



## CreateWebhookEndpoint

> WebhookEndpoint CreateWebhookEndpoint(ctx).WebhookEndpointCreateInput(webhookEndpointCreateInput).Execute()

Create a webhook_endpoint



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
    webhookEndpointCreateInput := *openapiclient.NewWebhookEndpointCreateInput() // WebhookEndpointCreateInput | Webhook Endpoint payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsAPI.CreateWebhookEndpoint(context.Background()).WebhookEndpointCreateInput(webhookEndpointCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.CreateWebhookEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhookEndpoint`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.CreateWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEndpointCreateInput** | [**WebhookEndpointCreateInput**](WebhookEndpointCreateInput.md) | Webhook Endpoint payload | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWebhookEndpoint

> WebhookEndpoint DestroyWebhookEndpoint(ctx, lagoId).Execute()

Delete a webhook endpoint



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsAPI.DestroyWebhookEndpoint(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.DestroyWebhookEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyWebhookEndpoint`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.DestroyWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllWebhookEndpoints

> WebhookEndpointsPaginated FindAllWebhookEndpoints(ctx).Page(page).PerPage(perPage).Execute()

List all webhook endpoints



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
    resp, r, err := apiClient.WebhookEndpointsAPI.FindAllWebhookEndpoints(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.FindAllWebhookEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllWebhookEndpoints`: WebhookEndpointsPaginated
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.FindAllWebhookEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllWebhookEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**WebhookEndpointsPaginated**](WebhookEndpointsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebhookEndpoint

> WebhookEndpoint FindWebhookEndpoint(ctx, lagoId).Execute()

Retrieve a webhook endpoint



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsAPI.FindWebhookEndpoint(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.FindWebhookEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWebhookEndpoint`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.FindWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookEndpoint

> WebhookEndpoint UpdateWebhookEndpoint(ctx, lagoId).WebhookEndpointUpdateInput(webhookEndpointUpdateInput).Execute()

Update a webhook endpoint



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.
    webhookEndpointUpdateInput := *openapiclient.NewWebhookEndpointUpdateInput() // WebhookEndpointUpdateInput | Webhook Endpoint update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsAPI.UpdateWebhookEndpoint(context.Background(), lagoId).WebhookEndpointUpdateInput(webhookEndpointUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.UpdateWebhookEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookEndpoint`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.UpdateWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookEndpointUpdateInput** | [**WebhookEndpointUpdateInput**](WebhookEndpointUpdateInput.md) | Webhook Endpoint update payload | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

