# \BillableMetricsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillableMetric**](BillableMetricsAPI.md#CreateBillableMetric) | **Post** /billable_metrics | Create a billable metric
[**DestroyBillableMetric**](BillableMetricsAPI.md#DestroyBillableMetric) | **Delete** /billable_metrics/{code} | Delete a billable metric
[**FindAllBillableMetricGroups**](BillableMetricsAPI.md#FindAllBillableMetricGroups) | **Get** /billable_metrics/{code}/groups | Find a billable metric&#39;s groups
[**FindAllBillableMetrics**](BillableMetricsAPI.md#FindAllBillableMetrics) | **Get** /billable_metrics | List all billable metrics
[**FindBillableMetric**](BillableMetricsAPI.md#FindBillableMetric) | **Get** /billable_metrics/{code} | Retrieve a billable metric
[**UpdateBillableMetric**](BillableMetricsAPI.md#UpdateBillableMetric) | **Put** /billable_metrics/{code} | Update a billable metric



## CreateBillableMetric

> BillableMetric CreateBillableMetric(ctx).BillableMetricCreateInput(billableMetricCreateInput).Execute()

Create a billable metric



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
    billableMetricCreateInput := *openapiclient.NewBillableMetricCreateInput(*openapiclient.NewBillableMetricCreateInputBillableMetric("Storage", "storage", "sum_agg")) // BillableMetricCreateInput | Billable metric payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillableMetricsAPI.CreateBillableMetric(context.Background()).BillableMetricCreateInput(billableMetricCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.CreateBillableMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillableMetric`: BillableMetric
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.CreateBillableMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billableMetricCreateInput** | [**BillableMetricCreateInput**](BillableMetricCreateInput.md) | Billable metric payload | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyBillableMetric

> BillableMetric DestroyBillableMetric(ctx, code).Execute()

Delete a billable metric



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
    code := "storage" // string | Code of the existing billable metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillableMetricsAPI.DestroyBillableMetric(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.DestroyBillableMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyBillableMetric`: BillableMetric
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.DestroyBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of the existing billable metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllBillableMetricGroups

> GroupsPaginated FindAllBillableMetricGroups(ctx, code).Page(page).PerPage(perPage).Execute()

Find a billable metric's groups



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
    code := "example_code" // string | Code of the existing billable metric.
    page := int32(1) // int32 | Page number. (optional)
    perPage := int32(20) // int32 | Number of records per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillableMetricsAPI.FindAllBillableMetricGroups(context.Background(), code).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.FindAllBillableMetricGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllBillableMetricGroups`: GroupsPaginated
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.FindAllBillableMetricGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of the existing billable metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllBillableMetricGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**GroupsPaginated**](GroupsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllBillableMetrics

> BillableMetricsPaginated FindAllBillableMetrics(ctx).Page(page).PerPage(perPage).Execute()

List all billable metrics



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
    resp, r, err := apiClient.BillableMetricsAPI.FindAllBillableMetrics(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.FindAllBillableMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllBillableMetrics`: BillableMetricsPaginated
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.FindAllBillableMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllBillableMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 

### Return type

[**BillableMetricsPaginated**](BillableMetricsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBillableMetric

> BillableMetric FindBillableMetric(ctx, code).Execute()

Retrieve a billable metric



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
    code := "storage" // string | Code of the existing billable metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillableMetricsAPI.FindBillableMetric(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.FindBillableMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindBillableMetric`: BillableMetric
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.FindBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of the existing billable metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillableMetric

> BillableMetric UpdateBillableMetric(ctx, code).BillableMetricUpdateInput(billableMetricUpdateInput).Execute()

Update a billable metric



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
    code := "storage" // string | Code of the existing billable metric.
    billableMetricUpdateInput := *openapiclient.NewBillableMetricUpdateInput(*openapiclient.NewBillableMetricBaseInput()) // BillableMetricUpdateInput | Billable metric payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillableMetricsAPI.UpdateBillableMetric(context.Background(), code).BillableMetricUpdateInput(billableMetricUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillableMetricsAPI.UpdateBillableMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBillableMetric`: BillableMetric
    fmt.Fprintf(os.Stdout, "Response from `BillableMetricsAPI.UpdateBillableMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of the existing billable metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillableMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billableMetricUpdateInput** | [**BillableMetricUpdateInput**](BillableMetricUpdateInput.md) | Billable metric payload | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

