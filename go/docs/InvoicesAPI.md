# \InvoicesAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](InvoicesAPI.md#CreateInvoice) | **Post** /invoices | Create a one-off invoice
[**DownloadInvoice**](InvoicesAPI.md#DownloadInvoice) | **Post** /invoices/{lago_id}/download | Download an invoice PDF
[**FinalizeInvoice**](InvoicesAPI.md#FinalizeInvoice) | **Put** /invoices/{lago_id}/finalize | Finalize a draft invoice
[**FindAllInvoices**](InvoicesAPI.md#FindAllInvoices) | **Get** /invoices | List all invoices
[**FindInvoice**](InvoicesAPI.md#FindInvoice) | **Get** /invoices/{lago_id} | Retrieve an invoice
[**RefreshInvoice**](InvoicesAPI.md#RefreshInvoice) | **Put** /invoices/{lago_id}/refresh | Refresh a draft invoice
[**RetryPayment**](InvoicesAPI.md#RetryPayment) | **Post** /invoices/{lago_id}/retry_payment | Retry an invoice payment
[**UpdateInvoice**](InvoicesAPI.md#UpdateInvoice) | **Put** /invoices/{lago_id} | Update an invoice



## CreateInvoice

> Invoice CreateInvoice(ctx).InvoiceOneOffCreateInput(invoiceOneOffCreateInput).Execute()

Create a one-off invoice



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
    invoiceOneOffCreateInput := *openapiclient.NewInvoiceOneOffCreateInput(*openapiclient.NewInvoiceOneOffCreateInputInvoice("hooli_1234", []openapiclient.InvoiceOneOffCreateInputInvoiceFeesInner{*openapiclient.NewInvoiceOneOffCreateInputInvoiceFeesInner("setup_fee")})) // InvoiceOneOffCreateInput | Invoice payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.CreateInvoice(context.Background()).InvoiceOneOffCreateInput(invoiceOneOffCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceOneOffCreateInput** | [**InvoiceOneOffCreateInput**](InvoiceOneOffCreateInput.md) | Invoice payload | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadInvoice

> Invoice DownloadInvoice(ctx, lagoId).Execute()

Download an invoice PDF



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.DownloadInvoice(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DownloadInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DownloadInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeInvoice

> Invoice FinalizeInvoice(ctx, lagoId).Execute()

Finalize a draft invoice



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.FinalizeInvoice(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.FinalizeInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.FinalizeInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInvoices

> InvoicesPaginated FindAllInvoices(ctx).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).IssuingDateFrom(issuingDateFrom).IssuingDateTo(issuingDateTo).Status(status).PaymentStatus(paymentStatus).Execute()

List all invoices



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
    issuingDateFrom := time.Now() // string | Filter invoices starting from a specific date. (optional)
    issuingDateTo := time.Now() // string | Filter invoices up to a specific date. (optional)
    status := "status_example" // string | Filter invoices by status. Possible values are `draft` or `finalized`. (optional)
    paymentStatus := "paymentStatus_example" // string | Filter invoices by payment status. Possible values are `pending`, `failed` or `succeeded`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.FindAllInvoices(context.Background()).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).IssuingDateFrom(issuingDateFrom).IssuingDateTo(issuingDateTo).Status(status).PaymentStatus(paymentStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.FindAllInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInvoices`: InvoicesPaginated
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.FindAllInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **externalCustomerId** | **string** | Unique identifier assigned to the customer in your application. | 
 **issuingDateFrom** | **string** | Filter invoices starting from a specific date. | 
 **issuingDateTo** | **string** | Filter invoices up to a specific date. | 
 **status** | **string** | Filter invoices by status. Possible values are &#x60;draft&#x60; or &#x60;finalized&#x60;. | 
 **paymentStatus** | **string** | Filter invoices by payment status. Possible values are &#x60;pending&#x60;, &#x60;failed&#x60; or &#x60;succeeded&#x60;. | 

### Return type

[**InvoicesPaginated**](InvoicesPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInvoice

> Invoice FindInvoice(ctx, lagoId).Execute()

Retrieve an invoice



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.FindInvoice(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.FindInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.FindInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshInvoice

> Invoice RefreshInvoice(ctx, lagoId).Execute()

Refresh a draft invoice



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.RefreshInvoice(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.RefreshInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.RefreshInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryPayment

> RetryPayment(ctx, lagoId).Execute()

Retry an invoice payment



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InvoicesAPI.RetryPayment(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.RetryPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> Invoice UpdateInvoice(ctx, lagoId).InvoiceUpdateInput(invoiceUpdateInput).Execute()

Update an invoice



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.
    invoiceUpdateInput := *openapiclient.NewInvoiceUpdateInput(*openapiclient.NewInvoiceUpdateInputInvoice()) // InvoiceUpdateInput | Update an existing invoice

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.UpdateInvoice(context.Background(), lagoId).InvoiceUpdateInput(invoiceUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceUpdateInput** | [**InvoiceUpdateInput**](InvoiceUpdateInput.md) | Update an existing invoice | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

