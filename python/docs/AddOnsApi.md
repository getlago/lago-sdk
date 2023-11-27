# lago_api.AddOnsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_add_on**](AddOnsApi.md#create_add_on) | **POST** /add_ons | Create an add-on
[**destroy_add_on**](AddOnsApi.md#destroy_add_on) | **DELETE** /add_ons/{code} | Delete an add-on
[**find_add_on**](AddOnsApi.md#find_add_on) | **GET** /add_ons/{code} | Retrieve an add-on
[**find_all_add_ons**](AddOnsApi.md#find_all_add_ons) | **GET** /add_ons | List all add-ons
[**update_add_on**](AddOnsApi.md#update_add_on) | **PUT** /add_ons/{code} | Update an add-on


# **create_add_on**
> AddOn create_add_on(add_on_create_input)

Create an add-on

This endpoint is used to create an add-on that can be then attached to a one-off invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.add_on import AddOn
from lago_api.models.add_on_create_input import AddOnCreateInput
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
    api_instance = lago_api.AddOnsApi(api_client)
    add_on_create_input = lago_api.AddOnCreateInput() # AddOnCreateInput | Add-on payload

    try:
        # Create an add-on
        api_response = api_instance.create_add_on(add_on_create_input)
        print("The response of AddOnsApi->create_add_on:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddOnsApi->create_add_on: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_on_create_input** | [**AddOnCreateInput**](AddOnCreateInput.md)| Add-on payload | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Add-on created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_add_on**
> AddOn destroy_add_on(code)

Delete an add-on

This endpoint is used to delete an existing add-on.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.add_on import AddOn
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
    api_instance = lago_api.AddOnsApi(api_client)
    code = 'setup_fee' # str | Unique code used to identify the add-on.

    try:
        # Delete an add-on
        api_response = api_instance.destroy_add_on(code)
        print("The response of AddOnsApi->destroy_add_on:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddOnsApi->destroy_add_on: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the add-on. | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Add-on deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_add_on**
> AddOn find_add_on(code)

Retrieve an add-on

This endpoint is used to retrieve a specific add-on.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.add_on import AddOn
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
    api_instance = lago_api.AddOnsApi(api_client)
    code = 'setup_fee' # str | Unique code used to identify the add-on.

    try:
        # Retrieve an add-on
        api_response = api_instance.find_add_on(code)
        print("The response of AddOnsApi->find_add_on:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddOnsApi->find_add_on: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the add-on. | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Add-on |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_add_ons**
> AddOnsPaginated find_all_add_ons(page=page, per_page=per_page)

List all add-ons

This endpoint is used to list all existing add-ons.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.add_ons_paginated import AddOnsPaginated
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
    api_instance = lago_api.AddOnsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all add-ons
        api_response = api_instance.find_all_add_ons(page=page, per_page=per_page)
        print("The response of AddOnsApi->find_all_add_ons:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddOnsApi->find_all_add_ons: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**AddOnsPaginated**](AddOnsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Add-ons |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_add_on**
> AddOn update_add_on(code, add_on_update_input)

Update an add-on

This endpoint is used to update an existing add-on.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.add_on import AddOn
from lago_api.models.add_on_update_input import AddOnUpdateInput
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
    api_instance = lago_api.AddOnsApi(api_client)
    code = 'setup_fee' # str | Unique code used to identify the add-on.
    add_on_update_input = lago_api.AddOnUpdateInput() # AddOnUpdateInput | Add-on payload

    try:
        # Update an add-on
        api_response = api_instance.update_add_on(code, add_on_update_input)
        print("The response of AddOnsApi->update_add_on:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddOnsApi->update_add_on: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the add-on. | 
 **add_on_update_input** | [**AddOnUpdateInput**](AddOnUpdateInput.md)| Add-on payload | 

### Return type

[**AddOn**](AddOn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Add-on updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

