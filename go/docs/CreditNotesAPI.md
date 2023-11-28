# \CreditNotesAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreditNote**](CreditNotesAPI.md#CreateCreditNote) | **Post** /credit_notes | Create a credit note
[**DownloadCreditNote**](CreditNotesAPI.md#DownloadCreditNote) | **Post** /credit_notes/{lago_id}/download | Download a credit note PDF
[**EstimateCreditNote**](CreditNotesAPI.md#EstimateCreditNote) | **Post** /credit_notes/estimate | Estimate amounts for a new credit note
[**FindAllCreditNotes**](CreditNotesAPI.md#FindAllCreditNotes) | **Get** /credit_notes | List all credit notes
[**FindCreditNote**](CreditNotesAPI.md#FindCreditNote) | **Get** /credit_notes/{lago_id} | Retrieve a credit note
[**UpdateCreditNote**](CreditNotesAPI.md#UpdateCreditNote) | **Put** /credit_notes/{lago_id} | Update a credit note
[**VoidCreditNote**](CreditNotesAPI.md#VoidCreditNote) | **Put** /credit_notes/{lago_id}/void | Void a credit note



## CreateCreditNote

> CreditNote CreateCreditNote(ctx).CreditNoteCreateInput(creditNoteCreateInput).Execute()

Create a credit note



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
    creditNoteCreateInput := *openapiclient.NewCreditNoteCreateInput(*openapiclient.NewCreditNoteCreateInputCreditNote("1a901a90-1a90-1a90-1a90-1a901a901a90", []openapiclient.CreditNoteEstimateInputCreditNoteItemsInner{*openapiclient.NewCreditNoteEstimateInputCreditNoteItemsInner("1a901a90-1a90-1a90-1a90-1a901a901a90", int32(10))})) // CreditNoteCreateInput | Credit note payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.CreateCreditNote(context.Background()).CreditNoteCreateInput(creditNoteCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreateCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreditNote`: CreditNote
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreateCreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteCreateInput** | [**CreditNoteCreateInput**](CreditNoteCreateInput.md) | Credit note payload | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCreditNote

> CreditNote DownloadCreditNote(ctx, lagoId).Execute()

Download a credit note PDF



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | The credit note unique identifier, created by Lago.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.DownloadCreditNote(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.DownloadCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadCreditNote`: CreditNote
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.DownloadCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | The credit note unique identifier, created by Lago. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateCreditNote

> CreditNoteEstimated EstimateCreditNote(ctx).CreditNoteEstimateInput(creditNoteEstimateInput).Execute()

Estimate amounts for a new credit note



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
    creditNoteEstimateInput := *openapiclient.NewCreditNoteEstimateInput(*openapiclient.NewCreditNoteEstimateInputCreditNote("1a901a90-1a90-1a90-1a90-1a901a901a90", []openapiclient.CreditNoteEstimateInputCreditNoteItemsInner{*openapiclient.NewCreditNoteEstimateInputCreditNoteItemsInner("1a901a90-1a90-1a90-1a90-1a901a901a90", int32(10))})) // CreditNoteEstimateInput | Credit note estimate payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.EstimateCreditNote(context.Background()).CreditNoteEstimateInput(creditNoteEstimateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.EstimateCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EstimateCreditNote`: CreditNoteEstimated
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.EstimateCreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteEstimateInput** | [**CreditNoteEstimateInput**](CreditNoteEstimateInput.md) | Credit note estimate payload | 

### Return type

[**CreditNoteEstimated**](CreditNoteEstimated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllCreditNotes

> CreditNotes FindAllCreditNotes(ctx).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).Execute()

List all credit notes



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
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | Unique identifier assigned to the customer in your application. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.FindAllCreditNotes(context.Background()).Page(page).PerPage(perPage).ExternalCustomerId(externalCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.FindAllCreditNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllCreditNotes`: CreditNotes
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.FindAllCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | 
 **perPage** | **int32** | Number of records per page. | 
 **externalCustomerId** | **string** | Unique identifier assigned to the customer in your application. | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCreditNote

> CreditNote FindCreditNote(ctx, lagoId).Execute()

Retrieve a credit note



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
    lagoId := "12345" // string | The credit note unique identifier, created by Lago.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.FindCreditNote(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.FindCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCreditNote`: CreditNote
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.FindCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | The credit note unique identifier, created by Lago. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNote

> CreditNote UpdateCreditNote(ctx, lagoId).CreditNoteUpdateInput(creditNoteUpdateInput).Execute()

Update a credit note



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
    lagoId := "12345" // string | The credit note unique identifier, created by Lago.
    creditNoteUpdateInput := *openapiclient.NewCreditNoteUpdateInput(*openapiclient.NewCreditNoteUpdateInputCreditNote("succeeded")) // CreditNoteUpdateInput | Credit note update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.UpdateCreditNote(context.Background(), lagoId).CreditNoteUpdateInput(creditNoteUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.UpdateCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreditNote`: CreditNote
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.UpdateCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | The credit note unique identifier, created by Lago. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditNoteUpdateInput** | [**CreditNoteUpdateInput**](CreditNoteUpdateInput.md) | Credit note update payload | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidCreditNote

> CreditNote VoidCreditNote(ctx, lagoId).Execute()

Void a credit note



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
    lagoId := "1a901a90-1a90-1a90-1a90-1a901a901a90" // string | The credit note unique identifier, created by Lago.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.VoidCreditNote(context.Background(), lagoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.VoidCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoidCreditNote`: CreditNote
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.VoidCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoId** | **string** | The credit note unique identifier, created by Lago. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

