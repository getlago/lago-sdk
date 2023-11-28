# \AnalyticsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAllGrossRevenues**](AnalyticsAPI.md#FindAllGrossRevenues) | **Get** /analytics/gross_revenue | List gross revenue
[**FindAllInvoiceCollections**](AnalyticsAPI.md#FindAllInvoiceCollections) | **Get** /analytics/invoice_collection | List of finalized invoices
[**FindAllInvoicedUsages**](AnalyticsAPI.md#FindAllInvoicedUsages) | **Get** /analytics/invoiced_usage | List usage revenue
[**FindAllMrrs**](AnalyticsAPI.md#FindAllMrrs) | **Get** /analytics/mrr | List MRR



## FindAllGrossRevenues

> GrossRevenues FindAllGrossRevenues(ctx).Currency(currency).ExternalCustomerId(externalCustomerId).Execute()

List gross revenue



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
    currency := openapiclient.Currency("AED") // Currency | Currency of revenue analytics. Format must be ISO 4217. (optional)
    externalCustomerId := "5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba" // string | The customer external unique identifier (provided by your own application). Use it to filter revenue analytics at the customer level. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsAPI.FindAllGrossRevenues(context.Background()).Currency(currency).ExternalCustomerId(externalCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FindAllGrossRevenues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllGrossRevenues`: GrossRevenues
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.FindAllGrossRevenues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllGrossRevenuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](Currency.md) | Currency of revenue analytics. Format must be ISO 4217. | 
 **externalCustomerId** | **string** | The customer external unique identifier (provided by your own application). Use it to filter revenue analytics at the customer level. | 

### Return type

[**GrossRevenues**](GrossRevenues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInvoiceCollections

> InvoiceCollections FindAllInvoiceCollections(ctx).Currency(currency).Execute()

List of finalized invoices



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
    currency := openapiclient.Currency("AED") // Currency | The currency of revenue analytics. Format must be ISO 4217. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsAPI.FindAllInvoiceCollections(context.Background()).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FindAllInvoiceCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInvoiceCollections`: InvoiceCollections
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.FindAllInvoiceCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInvoiceCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](Currency.md) | The currency of revenue analytics. Format must be ISO 4217. | 

### Return type

[**InvoiceCollections**](InvoiceCollections.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllInvoicedUsages

> InvoicedUsages FindAllInvoicedUsages(ctx).Currency(currency).Execute()

List usage revenue



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
    currency := openapiclient.Currency("AED") // Currency | The currency of invoiced usage analytics. Format must be ISO 4217. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsAPI.FindAllInvoicedUsages(context.Background()).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FindAllInvoicedUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllInvoicedUsages`: InvoicedUsages
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.FindAllInvoicedUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllInvoicedUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](Currency.md) | The currency of invoiced usage analytics. Format must be ISO 4217. | 

### Return type

[**InvoicedUsages**](InvoicedUsages.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllMrrs

> Mrrs FindAllMrrs(ctx).Currency(currency).Months(months).Execute()

List MRR



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
    currency := openapiclient.Currency("AED") // Currency | Quantifies the revenue generated from `subscription` fees on a monthly basis. This figure is calculated post-application of applicable taxes and deduction of any applicable discounts. The method of calculation varies based on the subscription billing cycle:  - Revenue from `monthly` subscription invoices is included in the MRR for the month in which the invoice is issued. - Revenue from `quarterly` subscription invoices is distributed evenly over three months. This distribution applies to fees paid in advance (allocated to the next remaining months depending on calendar or anniversary billing) as well as to fees paid in arrears (allocated to the preceding months depending on calendar or anniversary billing). - Revenue from `yearly` subscription invoices is distributed evenly over twelve months. This allocation is applicable for fees paid in advance (spread over the next  remaining months depending on calendar or anniversary billing) and for fees paid in arrears (spread over the previous months depending on calendar or anniversary billing). - Revenue from `weekly` subscription invoices, the total revenue from all invoices issued within a month is summed up. This total is then divided by the number of invoices issued during that month, and the result is multiplied by 4.33, representing the average number of weeks in a month. (optional)
    months := int32(12) // int32 | Show data only for given number of months. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsAPI.FindAllMrrs(context.Background()).Currency(currency).Months(months).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.FindAllMrrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllMrrs`: Mrrs
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.FindAllMrrs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAllMrrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](Currency.md) | Quantifies the revenue generated from &#x60;subscription&#x60; fees on a monthly basis. This figure is calculated post-application of applicable taxes and deduction of any applicable discounts. The method of calculation varies based on the subscription billing cycle:  - Revenue from &#x60;monthly&#x60; subscription invoices is included in the MRR for the month in which the invoice is issued. - Revenue from &#x60;quarterly&#x60; subscription invoices is distributed evenly over three months. This distribution applies to fees paid in advance (allocated to the next remaining months depending on calendar or anniversary billing) as well as to fees paid in arrears (allocated to the preceding months depending on calendar or anniversary billing). - Revenue from &#x60;yearly&#x60; subscription invoices is distributed evenly over twelve months. This allocation is applicable for fees paid in advance (spread over the next  remaining months depending on calendar or anniversary billing) and for fees paid in arrears (spread over the previous months depending on calendar or anniversary billing). - Revenue from &#x60;weekly&#x60; subscription invoices, the total revenue from all invoices issued within a month is summed up. This total is then divided by the number of invoices issued during that month, and the result is multiplied by 4.33, representing the average number of weeks in a month. | 
 **months** | **int32** | Show data only for given number of months. | 

### Return type

[**Mrrs**](Mrrs.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

