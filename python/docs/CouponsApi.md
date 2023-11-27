# lago_api.CouponsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apply_coupon**](CouponsApi.md#apply_coupon) | **POST** /applied_coupons | Apply a coupon to a customer
[**create_coupon**](CouponsApi.md#create_coupon) | **POST** /coupons | Create a coupon
[**delete_applied_coupon**](CouponsApi.md#delete_applied_coupon) | **DELETE** /customers/{external_customer_id}/applied_coupons/{applied_coupon_id} | Delete an applied coupon
[**destroy_coupon**](CouponsApi.md#destroy_coupon) | **DELETE** /coupons/{code} | Delete a coupon
[**find_all_applied_coupons**](CouponsApi.md#find_all_applied_coupons) | **GET** /applied_coupons | List all applied coupons
[**find_all_coupons**](CouponsApi.md#find_all_coupons) | **GET** /coupons | List all coupons
[**find_coupon**](CouponsApi.md#find_coupon) | **GET** /coupons/{code} | Retrieve a coupon
[**update_coupon**](CouponsApi.md#update_coupon) | **PUT** /coupons/{code} | Update a coupon


# **apply_coupon**
> AppliedCoupon apply_coupon(applied_coupon_input)

Apply a coupon to a customer

This endpoint is used to apply a specific coupon to a customer, before or during a subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.applied_coupon import AppliedCoupon
from lago_api.models.applied_coupon_input import AppliedCouponInput
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
    api_instance = lago_api.CouponsApi(api_client)
    applied_coupon_input = lago_api.AppliedCouponInput() # AppliedCouponInput | Apply coupon payload

    try:
        # Apply a coupon to a customer
        api_response = api_instance.apply_coupon(applied_coupon_input)
        print("The response of CouponsApi->apply_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->apply_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applied_coupon_input** | [**AppliedCouponInput**](AppliedCouponInput.md)| Apply coupon payload | 

### Return type

[**AppliedCoupon**](AppliedCoupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupon applied |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_coupon**
> Coupon create_coupon(coupon_create_input)

Create a coupon

This endpoint is used to create a coupon that can be then attached to a customer to create a discount.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.coupon import Coupon
from lago_api.models.coupon_create_input import CouponCreateInput
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
    api_instance = lago_api.CouponsApi(api_client)
    coupon_create_input = lago_api.CouponCreateInput() # CouponCreateInput | Coupon payload

    try:
        # Create a coupon
        api_response = api_instance.create_coupon(coupon_create_input)
        print("The response of CouponsApi->create_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->create_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coupon_create_input** | [**CouponCreateInput**](CouponCreateInput.md)| Coupon payload | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupon created |  -  |
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
    api_instance = lago_api.CouponsApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application)
    applied_coupon_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier of the applied coupon, created by Lago.

    try:
        # Delete an applied coupon
        api_response = api_instance.delete_applied_coupon(external_customer_id, applied_coupon_id)
        print("The response of CouponsApi->delete_applied_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->delete_applied_coupon: %s\n" % e)
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

# **destroy_coupon**
> Coupon destroy_coupon(code)

Delete a coupon

This endpoint is used to delete a coupon.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.coupon import Coupon
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
    api_instance = lago_api.CouponsApi(api_client)
    code = 'startup_deal' # str | Unique code used to identify the coupon.

    try:
        # Delete a coupon
        api_response = api_instance.destroy_coupon(code)
        print("The response of CouponsApi->destroy_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->destroy_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the coupon. | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupon deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_applied_coupons**
> AppliedCouponsPaginated find_all_applied_coupons(page=page, per_page=per_page, status=status, external_customer_id=external_customer_id)

List all applied coupons

This endpoint is used to list all applied coupons. You can filter by coupon status and by customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.applied_coupons_paginated import AppliedCouponsPaginated
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
    api_instance = lago_api.CouponsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    status = 'active' # str | The status of the coupon. Can be either `active` or `terminated`. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application) (optional)

    try:
        # List all applied coupons
        api_response = api_instance.find_all_applied_coupons(page=page, per_page=per_page, status=status, external_customer_id=external_customer_id)
        print("The response of CouponsApi->find_all_applied_coupons:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->find_all_applied_coupons: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **status** | **str**| The status of the coupon. Can be either &#x60;active&#x60; or &#x60;terminated&#x60;. | [optional] 
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application) | [optional] 

### Return type

[**AppliedCouponsPaginated**](AppliedCouponsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Applied Coupons |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_coupons**
> CouponsPaginated find_all_coupons(page=page, per_page=per_page)

List all coupons

This endpoint is used to list all existing coupons.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.coupons_paginated import CouponsPaginated
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
    api_instance = lago_api.CouponsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all coupons
        api_response = api_instance.find_all_coupons(page=page, per_page=per_page)
        print("The response of CouponsApi->find_all_coupons:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->find_all_coupons: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**CouponsPaginated**](CouponsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupons |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_coupon**
> Coupon find_coupon(code)

Retrieve a coupon

This endpoint is used to retrieve a specific coupon.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.coupon import Coupon
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
    api_instance = lago_api.CouponsApi(api_client)
    code = 'startup_deal' # str | Unique code used to identify the coupon.

    try:
        # Retrieve a coupon
        api_response = api_instance.find_coupon(code)
        print("The response of CouponsApi->find_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->find_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the coupon. | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupon |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_coupon**
> Coupon update_coupon(code, coupon_update_input)

Update a coupon

This endpoint is used to update a coupon that can be then attached to a customer to create a discount.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.coupon import Coupon
from lago_api.models.coupon_update_input import CouponUpdateInput
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
    api_instance = lago_api.CouponsApi(api_client)
    code = 'startup_deal' # str | Unique code used to identify the coupon.
    coupon_update_input = lago_api.CouponUpdateInput() # CouponUpdateInput | Coupon payload

    try:
        # Update a coupon
        api_response = api_instance.update_coupon(code, coupon_update_input)
        print("The response of CouponsApi->update_coupon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CouponsApi->update_coupon: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Unique code used to identify the coupon. | 
 **coupon_update_input** | [**CouponUpdateInput**](CouponUpdateInput.md)| Coupon payload | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Coupon updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

