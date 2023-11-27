# lago_api.WebhookEndpointsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_webhook_endpoint**](WebhookEndpointsApi.md#create_webhook_endpoint) | **POST** /webhook_endpoints | Create a webhook_endpoint
[**destroy_webhook_endpoint**](WebhookEndpointsApi.md#destroy_webhook_endpoint) | **DELETE** /webhook_endpoints/{lago_id} | Delete a webhook endpoint
[**find_all_webhook_endpoints**](WebhookEndpointsApi.md#find_all_webhook_endpoints) | **GET** /webhook_endpoints | List all webhook endpoints
[**find_webhook_endpoint**](WebhookEndpointsApi.md#find_webhook_endpoint) | **GET** /webhook_endpoints/{lago_id} | Retrieve a webhook endpoint
[**update_webhook_endpoint**](WebhookEndpointsApi.md#update_webhook_endpoint) | **PUT** /webhook_endpoints/{lago_id} | Update a webhook endpoint


# **create_webhook_endpoint**
> WebhookEndpoint create_webhook_endpoint(webhook_endpoint_create_input)

Create a webhook_endpoint

This endpoint is used to create a webhook endpoint.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.webhook_endpoint import WebhookEndpoint
from lago_api.models.webhook_endpoint_create_input import WebhookEndpointCreateInput
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
    api_instance = lago_api.WebhookEndpointsApi(api_client)
    webhook_endpoint_create_input = lago_api.WebhookEndpointCreateInput() # WebhookEndpointCreateInput | Webhook Endpoint payload

    try:
        # Create a webhook_endpoint
        api_response = api_instance.create_webhook_endpoint(webhook_endpoint_create_input)
        print("The response of WebhookEndpointsApi->create_webhook_endpoint:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WebhookEndpointsApi->create_webhook_endpoint: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhook_endpoint_create_input** | [**WebhookEndpointCreateInput**](WebhookEndpointCreateInput.md)| Webhook Endpoint payload | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Webhook Endpoint created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_webhook_endpoint**
> WebhookEndpoint destroy_webhook_endpoint(lago_id)

Delete a webhook endpoint

This endpoint is used to delete an existing webhook endpoint.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.webhook_endpoint import WebhookEndpoint
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
    api_instance = lago_api.WebhookEndpointsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.

    try:
        # Delete a webhook endpoint
        api_response = api_instance.destroy_webhook_endpoint(lago_id)
        print("The response of WebhookEndpointsApi->destroy_webhook_endpoint:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WebhookEndpointsApi->destroy_webhook_endpoint: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Webhook Endpoint deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**405** | Not Allowed error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_webhook_endpoints**
> WebhookEndpointsPaginated find_all_webhook_endpoints(page=page, per_page=per_page)

List all webhook endpoints

This endpoint is used to list all webhook endpoints.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.webhook_endpoints_paginated import WebhookEndpointsPaginated
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
    api_instance = lago_api.WebhookEndpointsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all webhook endpoints
        api_response = api_instance.find_all_webhook_endpoints(page=page, per_page=per_page)
        print("The response of WebhookEndpointsApi->find_all_webhook_endpoints:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WebhookEndpointsApi->find_all_webhook_endpoints: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**WebhookEndpointsPaginated**](WebhookEndpointsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | WebhookEndpoints |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_webhook_endpoint**
> WebhookEndpoint find_webhook_endpoint(lago_id)

Retrieve a webhook endpoint

This endpoint is used to retrieve an existing webhook endpoint.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.webhook_endpoint import WebhookEndpoint
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
    api_instance = lago_api.WebhookEndpointsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.

    try:
        # Retrieve a webhook endpoint
        api_response = api_instance.find_webhook_endpoint(lago_id)
        print("The response of WebhookEndpointsApi->find_webhook_endpoint:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WebhookEndpointsApi->find_webhook_endpoint: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | WebhookEndpoint |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_webhook_endpoint**
> WebhookEndpoint update_webhook_endpoint(lago_id, webhook_endpoint_update_input)

Update a webhook endpoint

This endpoint is used to update an existing webhook endpoint.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.webhook_endpoint import WebhookEndpoint
from lago_api.models.webhook_endpoint_update_input import WebhookEndpointUpdateInput
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
    api_instance = lago_api.WebhookEndpointsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint's record within the Lago system.
    webhook_endpoint_update_input = lago_api.WebhookEndpointUpdateInput() # WebhookEndpointUpdateInput | Webhook Endpoint update payload

    try:
        # Update a webhook endpoint
        api_response = api_instance.update_webhook_endpoint(lago_id, webhook_endpoint_update_input)
        print("The response of WebhookEndpointsApi->update_webhook_endpoint:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WebhookEndpointsApi->update_webhook_endpoint: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the webhook endpoint within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the webhook endpoint&#39;s record within the Lago system. | 
 **webhook_endpoint_update_input** | [**WebhookEndpointUpdateInput**](WebhookEndpointUpdateInput.md)| Webhook Endpoint update payload | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Webhook Endpoint updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

