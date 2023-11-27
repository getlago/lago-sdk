# lago_api.TaxesApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_tax**](TaxesApi.md#create_tax) | **POST** /taxes | Create a tax
[**destroy_tax**](TaxesApi.md#destroy_tax) | **DELETE** /taxes/{code} | Delete a tax
[**find_all_taxes**](TaxesApi.md#find_all_taxes) | **GET** /taxes | List all taxes
[**find_tax**](TaxesApi.md#find_tax) | **GET** /taxes/{code} | Retrieve a Tax
[**update_tax**](TaxesApi.md#update_tax) | **PUT** /taxes/{code} | Update a tax


# **create_tax**
> Tax create_tax(tax_create_input)

Create a tax

This endpoint creates a new tax representing a customizable tax rate applicable to either the organization or a specific customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.tax import Tax
from lago_api.models.tax_create_input import TaxCreateInput
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
    api_instance = lago_api.TaxesApi(api_client)
    tax_create_input = lago_api.TaxCreateInput() # TaxCreateInput | Tax creation payload

    try:
        # Create a tax
        api_response = api_instance.create_tax(tax_create_input)
        print("The response of TaxesApi->create_tax:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxesApi->create_tax: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_create_input** | [**TaxCreateInput**](TaxCreateInput.md)| Tax creation payload | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_tax**
> Tax destroy_tax(code)

Delete a tax

This endpoint is used to delete a tax.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.tax import Tax
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
    api_instance = lago_api.TaxesApi(api_client)
    code = 'french_standard_vat' # str | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.

    try:
        # Delete a tax
        api_response = api_instance.destroy_tax(code)
        print("The response of TaxesApi->destroy_tax:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxesApi->destroy_tax: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_taxes**
> TaxesPaginated find_all_taxes(page=page, per_page=per_page)

List all taxes

This endpoint retrieves all existing taxes representing a customizable tax rate applicable to either the organization or a specific customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.taxes_paginated import TaxesPaginated
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
    api_instance = lago_api.TaxesApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all taxes
        api_response = api_instance.find_all_taxes(page=page, per_page=per_page)
        print("The response of TaxesApi->find_all_taxes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxesApi->find_all_taxes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**TaxesPaginated**](TaxesPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Taxes |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_tax**
> Tax find_tax(code)

Retrieve a Tax

This endpoint retrieves an existing tax representing a customizable tax rate applicable to either the organization or a specific customer. The tax is identified by its unique code.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.tax import Tax
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
    api_instance = lago_api.TaxesApi(api_client)
    code = 'french_standard_vat' # str | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.

    try:
        # Retrieve a Tax
        api_response = api_instance.find_tax(code)
        print("The response of TaxesApi->find_tax:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxesApi->find_tax: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_tax**
> Tax update_tax(code, tax_update_input)

Update a tax

This endpoint updates an existing tax representing a customizable tax rate applicable to either the organization or a specific customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.tax import Tax
from lago_api.models.tax_update_input import TaxUpdateInput
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
    api_instance = lago_api.TaxesApi(api_client)
    code = 'french_standard_vat' # str | The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes.
    tax_update_input = lago_api.TaxUpdateInput() # TaxUpdateInput | Tax update payload

    try:
        # Update a tax
        api_response = api_instance.update_tax(code, tax_update_input)
        print("The response of TaxesApi->update_tax:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxesApi->update_tax: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the tax. It serves as a unique identifier associated with a particular tax. The code is typically used for internal or system-level identification purposes. | 
 **tax_update_input** | [**TaxUpdateInput**](TaxUpdateInput.md)| Tax update payload | 

### Return type

[**Tax**](Tax.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

