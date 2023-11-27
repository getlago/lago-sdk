# lago_api.InvoicesApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_invoice**](InvoicesApi.md#create_invoice) | **POST** /invoices | Create a one-off invoice
[**download_invoice**](InvoicesApi.md#download_invoice) | **POST** /invoices/{lago_id}/download | Download an invoice PDF
[**finalize_invoice**](InvoicesApi.md#finalize_invoice) | **PUT** /invoices/{lago_id}/finalize | Finalize a draft invoice
[**find_all_invoices**](InvoicesApi.md#find_all_invoices) | **GET** /invoices | List all invoices
[**find_invoice**](InvoicesApi.md#find_invoice) | **GET** /invoices/{lago_id} | Retrieve an invoice
[**refresh_invoice**](InvoicesApi.md#refresh_invoice) | **PUT** /invoices/{lago_id}/refresh | Refresh a draft invoice
[**retry_payment**](InvoicesApi.md#retry_payment) | **POST** /invoices/{lago_id}/retry_payment | Retry an invoice payment
[**update_invoice**](InvoicesApi.md#update_invoice) | **PUT** /invoices/{lago_id} | Update an invoice


# **create_invoice**
> Invoice create_invoice(invoice_one_off_create_input)

Create a one-off invoice

This endpoint is used for issuing a one-off invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
from lago_api.models.invoice_one_off_create_input import InvoiceOneOffCreateInput
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
    api_instance = lago_api.InvoicesApi(api_client)
    invoice_one_off_create_input = lago_api.InvoiceOneOffCreateInput() # InvoiceOneOffCreateInput | Invoice payload

    try:
        # Create a one-off invoice
        api_response = api_instance.create_invoice(invoice_one_off_create_input)
        print("The response of InvoicesApi->create_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->create_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice_one_off_create_input** | [**InvoiceOneOffCreateInput**](InvoiceOneOffCreateInput.md)| Invoice payload | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_invoice**
> Invoice download_invoice(lago_id)

Download an invoice PDF

This endpoint is used for downloading a specific invoice PDF document.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    try:
        # Download an invoice PDF
        api_response = api_instance.download_invoice(lago_id)
        print("The response of InvoicesApi->download_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->download_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice PDF |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **finalize_invoice**
> Invoice finalize_invoice(lago_id)

Finalize a draft invoice

This endpoint is used for finalizing a draft invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    try:
        # Finalize a draft invoice
        api_response = api_instance.finalize_invoice(lago_id)
        print("The response of InvoicesApi->finalize_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->finalize_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice finalized |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_invoices**
> InvoicesPaginated find_all_invoices(page=page, per_page=per_page, external_customer_id=external_customer_id, issuing_date_from=issuing_date_from, issuing_date_to=issuing_date_to, status=status, payment_status=payment_status)

List all invoices

This endpoint is used for retrievign all invoices.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoices_paginated import InvoicesPaginated
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
    api_instance = lago_api.InvoicesApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | Unique identifier assigned to the customer in your application. (optional)
    issuing_date_from = 'Fri Jul 08 00:00:00 UTC 2022' # date | Filter invoices starting from a specific date. (optional)
    issuing_date_to = 'Tue Aug 09 00:00:00 UTC 2022' # date | Filter invoices up to a specific date. (optional)
    status = 'status_example' # str | Filter invoices by status. Possible values are `draft` or `finalized`. (optional)
    payment_status = 'payment_status_example' # str | Filter invoices by payment status. Possible values are `pending`, `failed` or `succeeded`. (optional)

    try:
        # List all invoices
        api_response = api_instance.find_all_invoices(page=page, per_page=per_page, external_customer_id=external_customer_id, issuing_date_from=issuing_date_from, issuing_date_to=issuing_date_to, status=status, payment_status=payment_status)
        print("The response of InvoicesApi->find_all_invoices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->find_all_invoices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **external_customer_id** | **str**| Unique identifier assigned to the customer in your application. | [optional] 
 **issuing_date_from** | **date**| Filter invoices starting from a specific date. | [optional] 
 **issuing_date_to** | **date**| Filter invoices up to a specific date. | [optional] 
 **status** | **str**| Filter invoices by status. Possible values are &#x60;draft&#x60; or &#x60;finalized&#x60;. | [optional] 
 **payment_status** | **str**| Filter invoices by payment status. Possible values are &#x60;pending&#x60;, &#x60;failed&#x60; or &#x60;succeeded&#x60;. | [optional] 

### Return type

[**InvoicesPaginated**](InvoicesPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoices |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_invoice**
> Invoice find_invoice(lago_id)

Retrieve an invoice

This endpoint is used for retrieving a specific invoice that has been issued.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    try:
        # Retrieve an invoice
        api_response = api_instance.find_invoice(lago_id)
        print("The response of InvoicesApi->find_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->find_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **refresh_invoice**
> Invoice refresh_invoice(lago_id)

Refresh a draft invoice

This endpoint is used for refreshing a draft invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    try:
        # Refresh a draft invoice
        api_response = api_instance.refresh_invoice(lago_id)
        print("The response of InvoicesApi->refresh_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->refresh_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice refreshed |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **retry_payment**
> retry_payment(lago_id)

Retry an invoice payment

This endpoint resends an invoice for collection and retry a payment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.

    try:
        # Retry an invoice payment
        api_instance.retry_payment(lago_id)
    except Exception as e:
        print("Exception when calling InvoicesApi->retry_payment: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice payment retried |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**405** | Not Allowed error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_invoice**
> Invoice update_invoice(lago_id, invoice_update_input)

Update an invoice

This endpoint is used for updating an existing invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.invoice import Invoice
from lago_api.models.invoice_update_input import InvoiceUpdateInput
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
    api_instance = lago_api.InvoicesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system.
    invoice_update_input = lago_api.InvoiceUpdateInput() # InvoiceUpdateInput | Update an existing invoice

    try:
        # Update an invoice
        api_response = api_instance.update_invoice(lago_id, invoice_update_input)
        print("The response of InvoicesApi->update_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoicesApi->update_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the invoice within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the invoice’s record within the Lago system. | 
 **invoice_update_input** | [**InvoiceUpdateInput**](InvoiceUpdateInput.md)| Update an existing invoice | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

