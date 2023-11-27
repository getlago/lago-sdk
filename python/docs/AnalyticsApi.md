# lago_api.AnalyticsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**find_all_gross_revenues**](AnalyticsApi.md#find_all_gross_revenues) | **GET** /analytics/gross_revenue | List gross revenue
[**find_all_invoice_collections**](AnalyticsApi.md#find_all_invoice_collections) | **GET** /analytics/invoice_collection | List of finalized invoices
[**find_all_invoiced_usages**](AnalyticsApi.md#find_all_invoiced_usages) | **GET** /analytics/invoiced_usage | List usage revenue
[**find_all_mrrs**](AnalyticsApi.md#find_all_mrrs) | **GET** /analytics/mrr | List MRR


# **find_all_gross_revenues**
> GrossRevenues find_all_gross_revenues(currency=currency, external_customer_id=external_customer_id)

List gross revenue

Gross revenue is the sum of monthly `finalized` invoice payments and fees paid in advance that are not invoiceable. This total is calculated after deducting taxes and discounts.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.gross_revenues import GrossRevenues
from lago_api.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.getlago.com/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = lago_api.Configuration(
    host = "https://api.getlago.com/api/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = lago_api.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with lago_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = lago_api.AnalyticsApi(api_client)
    currency = lago_api.Currency() # Currency | Currency of revenue analytics. Format must be ISO 4217. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application). Use it to filter revenue analytics at the customer level. (optional)

    try:
        # List gross revenue
        api_response = api_instance.find_all_gross_revenues(currency=currency, external_customer_id=external_customer_id)
        print("The response of AnalyticsApi->find_all_gross_revenues:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AnalyticsApi->find_all_gross_revenues: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](.md)| Currency of revenue analytics. Format must be ISO 4217. | [optional] 
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application). Use it to filter revenue analytics at the customer level. | [optional] 

### Return type

[**GrossRevenues**](GrossRevenues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Gross revenue |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_invoice_collections**
> InvoiceCollections find_all_invoice_collections(currency=currency)

List of finalized invoices

Represents a monthly aggregation, detailing both the total count and the cumulative amount of invoices that have been marked as `finalized`. This report sorts invoices categorically based on their `payment_status`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice_collections import InvoiceCollections
from lago_api.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.getlago.com/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = lago_api.Configuration(
    host = "https://api.getlago.com/api/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = lago_api.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with lago_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = lago_api.AnalyticsApi(api_client)
    currency = lago_api.Currency() # Currency | The currency of revenue analytics. Format must be ISO 4217. (optional)

    try:
        # List of finalized invoices
        api_response = api_instance.find_all_invoice_collections(currency=currency)
        print("The response of AnalyticsApi->find_all_invoice_collections:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AnalyticsApi->find_all_invoice_collections: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](.md)| The currency of revenue analytics. Format must be ISO 4217. | [optional] 

### Return type

[**InvoiceCollections**](InvoiceCollections.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Finalized invoice collection |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_invoiced_usages**
> InvoicedUsages find_all_invoiced_usages(currency=currency)

List usage revenue

Reports a monthly analysis focused on the revenue generated from all usage-based fees. It exclusively accounts for revenue that has been formally invoiced. Importantly, this report does not include revenue related to the usage in the current billing period, limiting its scope to previously invoiced amounts.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoiced_usages import InvoicedUsages
from lago_api.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.getlago.com/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = lago_api.Configuration(
    host = "https://api.getlago.com/api/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = lago_api.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with lago_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = lago_api.AnalyticsApi(api_client)
    currency = lago_api.Currency() # Currency | The currency of invoiced usage analytics. Format must be ISO 4217. (optional)

    try:
        # List usage revenue
        api_response = api_instance.find_all_invoiced_usages(currency=currency)
        print("The response of AnalyticsApi->find_all_invoiced_usages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AnalyticsApi->find_all_invoiced_usages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](.md)| The currency of invoiced usage analytics. Format must be ISO 4217. | [optional] 

### Return type

[**InvoicedUsages**](InvoicedUsages.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoiced usage |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_mrrs**
> Mrrs find_all_mrrs(currency=currency, months=months)

List MRR

This endpoint is used to list MRR.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.mrrs import Mrrs
from lago_api.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.getlago.com/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = lago_api.Configuration(
    host = "https://api.getlago.com/api/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = lago_api.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with lago_api.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = lago_api.AnalyticsApi(api_client)
    currency = lago_api.Currency() # Currency | Quantifies the revenue generated from `subscription` fees on a monthly basis. This figure is calculated post-application of applicable taxes and prior to the deduction of any applicable discounts. The method of calculation varies based on the subscription billing cycle:  - Revenue from `monthly` subscription invoices is included in the MRR for the month in which the invoice is issued. - Revenue from `quarterly` subscription invoices is distributed evenly over three months. This distribution applies to fees paid in advance (allocated to the next three months) as well as to fees paid in arrears (allocated to the preceding three months). - Revenue from `yearly` subscription invoices is distributed evenly over twelve months. This allocation is applicable for fees paid in advance (spread over the next twelve months) and for fees paid in arrears (spread over the previous twelve months). - Revenue from `weekly` subscription invoices, the total revenue from all invoices issued within a month is summed up. This total is then divided by the number of invoices issued during that month, and the result is multiplied by 4.33, representing the average number of weeks in a month. (optional)
    months = 12 # int | Show data only for given number of months. (optional)

    try:
        # List MRR
        api_response = api_instance.find_all_mrrs(currency=currency, months=months)
        print("The response of AnalyticsApi->find_all_mrrs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AnalyticsApi->find_all_mrrs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | [**Currency**](.md)| Quantifies the revenue generated from &#x60;subscription&#x60; fees on a monthly basis. This figure is calculated post-application of applicable taxes and prior to the deduction of any applicable discounts. The method of calculation varies based on the subscription billing cycle:  - Revenue from &#x60;monthly&#x60; subscription invoices is included in the MRR for the month in which the invoice is issued. - Revenue from &#x60;quarterly&#x60; subscription invoices is distributed evenly over three months. This distribution applies to fees paid in advance (allocated to the next three months) as well as to fees paid in arrears (allocated to the preceding three months). - Revenue from &#x60;yearly&#x60; subscription invoices is distributed evenly over twelve months. This allocation is applicable for fees paid in advance (spread over the next twelve months) and for fees paid in arrears (spread over the previous twelve months). - Revenue from &#x60;weekly&#x60; subscription invoices, the total revenue from all invoices issued within a month is summed up. This total is then divided by the number of invoices issued during that month, and the result is multiplied by 4.33, representing the average number of weeks in a month. | [optional] 
 **months** | **int**| Show data only for given number of months. | [optional] 

### Return type

[**Mrrs**](Mrrs.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MRR |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

