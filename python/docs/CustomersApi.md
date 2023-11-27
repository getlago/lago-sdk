# lago_api.CustomersApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_customer**](CustomersApi.md#create_customer) | **POST** /customers | Create a customer
[**delete_applied_coupon**](CustomersApi.md#delete_applied_coupon) | **DELETE** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
[**destroy_customer**](CustomersApi.md#destroy_customer) | **DELETE** /customers/{external_id} | Delete a customer
[**find_all_customer_past_usage**](CustomersApi.md#find_all_customer_past_usage) | **GET** /customers/{external_customer_id}/past_usage | Retrieve customer past usage
[**find_all_customers**](CustomersApi.md#find_all_customers) | **GET** /customers | List all customers
[**find_customer**](CustomersApi.md#find_customer) | **GET** /customers/{external_id} | Retrieve a customer
[**find_customer_current_usage**](CustomersApi.md#find_customer_current_usage) | **GET** /customers/{external_customer_id}/current_usage | Retrieve customer current usage
[**generate_customer_checkout_url**](CustomersApi.md#generate_customer_checkout_url) | **POST** /customers/{external_customer_id}/checkout_url | Generate a Customer Payment Provider Checkout URL
[**get_customer_portal_url**](CustomersApi.md#get_customer_portal_url) | **GET** /customers/{external_customer_id}/portal_url | Get a customer portal URL


# **create_customer**
> Customer create_customer(customer_create_input)

Create a customer

This endpoint creates a new customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customer import Customer
from lago_api.models.customer_create_input import CustomerCreateInput
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
    api_instance = lago_api.CustomersApi(api_client)
    customer_create_input = lago_api.CustomerCreateInput() # CustomerCreateInput | Customer payload

    try:
        # Create a customer
        api_response = api_instance.create_customer(customer_create_input)
        print("The response of CustomersApi->create_customer:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->create_customer: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer_create_input** | [**CustomerCreateInput**](CustomerCreateInput.md)| Customer payload | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer created or updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_applied_coupon**
> AppliedCoupon delete_applied_coupon(external_customer_id, applied_coupon_id)

Delete an applied coupon

This endpoint is used to delete a specific coupon that has been applied to a customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.applied_coupon import AppliedCoupon
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
    api_instance = lago_api.CustomersApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application)
    applied_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier of the applied coupon, created by Lago.

    try:
        # Delete an applied coupon
        api_response = api_instance.delete_applied_coupon(external_customer_id, applied_coupon_id)
        print("The response of CustomersApi->delete_applied_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->delete_applied_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application) | 
 **applied_coupon_id** | **str**| Unique identifier of the applied coupon, created by Lago. | 

### Return type

[**AppliedCoupon**](AppliedCoupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_customer**
> Customer destroy_customer(external_id)

Delete a customer

This endpoint deletes an existing customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customer import Customer
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
    api_instance = lago_api.CustomersApi(api_client)
    external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing customer

    try:
        # Delete a customer
        api_response = api_instance.destroy_customer(external_id)
        print("The response of CustomersApi->destroy_customer:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->destroy_customer: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_id** | **str**| External ID of the existing customer | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_customer_past_usage**
> CustomerPastUsage find_all_customer_past_usage(external_customer_id, external_subscription_id, page=page, per_page=per_page, billable_metric_code=billable_metric_code, periods_count=periods_count)

Retrieve customer past usage

This endpoint enables the retrieval of the usage-based billing data for a customer within past periods.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customer_past_usage import CustomerPastUsage
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
    api_instance = lago_api.CustomersApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application).
    external_subscription_id = 'sub_1234567890' # str | The unique identifier of the subscription within your application.
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    billable_metric_code = 'cpu' # str | Billable metric code filter to apply to the charge usage (optional)
    periods_count = 5 # int | Number of past billing period to returns in the result (optional)

    try:
        # Retrieve customer past usage
        api_response = api_instance.find_all_customer_past_usage(external_customer_id, external_subscription_id, page=page, per_page=per_page, billable_metric_code=billable_metric_code, periods_count=periods_count)
        print("The response of CustomersApi->find_all_customer_past_usage:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->find_all_customer_past_usage: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application). | 
 **external_subscription_id** | **str**| The unique identifier of the subscription within your application. | 
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **billable_metric_code** | **str**| Billable metric code filter to apply to the charge usage | [optional] 
 **periods_count** | **int**| Number of past billing period to returns in the result | [optional] 

### Return type

[**CustomerPastUsage**](CustomerPastUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer past usage |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_customers**
> CustomersPaginated find_all_customers(page=page, per_page=per_page)

List all customers

This endpoint retrieves all existing customers.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customers_paginated import CustomersPaginated
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
    api_instance = lago_api.CustomersApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all customers
        api_response = api_instance.find_all_customers(page=page, per_page=per_page)
        print("The response of CustomersApi->find_all_customers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->find_all_customers: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**CustomersPaginated**](CustomersPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of customers |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_customer**
> Customer find_customer(external_id)

Retrieve a customer

This endpoint retrieves an existing customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customer import Customer
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
    api_instance = lago_api.CustomersApi(api_client)
    external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing customer

    try:
        # Retrieve a customer
        api_response = api_instance.find_customer(external_id)
        print("The response of CustomersApi->find_customer:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->find_customer: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_id** | **str**| External ID of the existing customer | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_customer_current_usage**
> CustomerUsage find_customer_current_usage(external_customer_id, external_subscription_id)

Retrieve customer current usage

This endpoint enables the retrieval of the usage-based billing data for a customer within the current period.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.customer_usage import CustomerUsage
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
    api_instance = lago_api.CustomersApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application).
    external_subscription_id = 'sub_1234567890' # str | The unique identifier of the subscription within your application.

    try:
        # Retrieve customer current usage
        api_response = api_instance.find_customer_current_usage(external_customer_id, external_subscription_id)
        print("The response of CustomersApi->find_customer_current_usage:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->find_customer_current_usage: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application). | 
 **external_subscription_id** | **str**| The unique identifier of the subscription within your application. | 

### Return type

[**CustomerUsage**](CustomerUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer usage |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **generate_customer_checkout_url**
> GenerateCustomerCheckoutURL200Response generate_customer_checkout_url(external_customer_id)

Generate a Customer Payment Provider Checkout URL

This endpoint regenerates the Payment Provider Checkout URL of a Customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.generate_customer_checkout_url200_response import GenerateCustomerCheckoutURL200Response
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
    api_instance = lago_api.CustomersApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application).

    try:
        # Generate a Customer Payment Provider Checkout URL
        api_response = api_instance.generate_customer_checkout_url(external_customer_id)
        print("The response of CustomersApi->generate_customer_checkout_url:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->generate_customer_checkout_url: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application). | 

### Return type

[**GenerateCustomerCheckoutURL200Response**](GenerateCustomerCheckoutURL200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Customer Checkout URL |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_customer_portal_url**
> GetCustomerPortalUrl200Response get_customer_portal_url(external_customer_id)

Get a customer portal URL

Retrieves an embeddable link for displaying a customer portal.  This endpoint allows you to fetch the URL that can be embedded to provide customers access to a dedicated portal

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.get_customer_portal_url200_response import GetCustomerPortalUrl200Response
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
    api_instance = lago_api.CustomersApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing customer

    try:
        # Get a customer portal URL
        api_response = api_instance.get_customer_portal_url(external_customer_id)
        print("The response of CustomersApi->get_customer_portal_url:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CustomersApi->get_customer_portal_url: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| External ID of the existing customer | 

### Return type

[**GetCustomerPortalUrl200Response**](GetCustomerPortalUrl200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Portal URL |  -  |
**401** | Unauthorized error |  -  |
**403** | Forbidden |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

