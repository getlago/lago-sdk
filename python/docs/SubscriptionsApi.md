# lago_api.SubscriptionsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_subscription**](SubscriptionsApi.md#create_subscription) | **POST** /subscriptions | Assign a plan to a customer
[**destroy_subscription**](SubscriptionsApi.md#destroy_subscription) | **DELETE** /subscriptions/{external_id} | Terminate a subscription
[**find_all_subscriptions**](SubscriptionsApi.md#find_all_subscriptions) | **GET** /subscriptions | List all subscriptions
[**find_subscription**](SubscriptionsApi.md#find_subscription) | **GET** /subscriptions/{external_id} | Retrieve a subscription
[**update_subscription**](SubscriptionsApi.md#update_subscription) | **PUT** /subscriptions/{external_id} | Update a subscription


# **create_subscription**
> Subscription create_subscription(subscription_create_input)

Assign a plan to a customer

This endpoint assigns a plan to a customer, creating or modifying a subscription. Ideal for initial subscriptions or plan changes (upgrades/downgrades).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.subscription import Subscription
from lago_api.models.subscription_create_input import SubscriptionCreateInput
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
    api_instance = lago_api.SubscriptionsApi(api_client)
    subscription_create_input = lago_api.SubscriptionCreateInput() # SubscriptionCreateInput | Subscription payload

    try:
        # Assign a plan to a customer
        api_response = api_instance.create_subscription(subscription_create_input)
        print("The response of SubscriptionsApi->create_subscription:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubscriptionsApi->create_subscription: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_create_input** | [**SubscriptionCreateInput**](SubscriptionCreateInput.md)| Subscription payload | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Subscription created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_subscription**
> Subscription destroy_subscription(external_id, status=status)

Terminate a subscription

This endpoint allows you to terminate a subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.subscription import Subscription
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
    api_instance = lago_api.SubscriptionsApi(api_client)
    external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing subscription
    status = 'pending' # str | If the field is not defined, Lago will terminate only `active` subscriptions. However, if you wish to cancel a `pending` subscription, please ensure that you include `status=pending` in your request. (optional)

    try:
        # Terminate a subscription
        api_response = api_instance.destroy_subscription(external_id, status=status)
        print("The response of SubscriptionsApi->destroy_subscription:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubscriptionsApi->destroy_subscription: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_id** | **str**| External ID of the existing subscription | 
 **status** | **str**| If the field is not defined, Lago will terminate only &#x60;active&#x60; subscriptions. However, if you wish to cancel a &#x60;pending&#x60; subscription, please ensure that you include &#x60;status&#x3D;pending&#x60; in your request. | [optional] 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Subscription terminated |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**405** | Not Allowed error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_subscriptions**
> SubscriptionsPaginated find_all_subscriptions(page=page, per_page=per_page, external_customer_id=external_customer_id, plan_code=plan_code, status=status)

List all subscriptions

This endpoint retrieves all active subscriptions.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.subscriptions_paginated import SubscriptionsPaginated
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
    api_instance = lago_api.SubscriptionsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application) (optional)
    plan_code = 'premium' # str | The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans. (optional)
    status = ['[\"active\",\"pending\"]'] # List[str] | If the field is not defined, Lago will return only `active` subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: `pending`, `canceled`, `terminated`, `active`. (optional)

    try:
        # List all subscriptions
        api_response = api_instance.find_all_subscriptions(page=page, per_page=per_page, external_customer_id=external_customer_id, plan_code=plan_code, status=status)
        print("The response of SubscriptionsApi->find_all_subscriptions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubscriptionsApi->find_all_subscriptions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application) | [optional] 
 **plan_code** | **str**| The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans. | [optional] 
 **status** | [**List[str]**](str.md)| If the field is not defined, Lago will return only &#x60;active&#x60; subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: &#x60;pending&#x60;, &#x60;canceled&#x60;, &#x60;terminated&#x60;, &#x60;active&#x60;. | [optional] 

### Return type

[**SubscriptionsPaginated**](SubscriptionsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of subscriptions |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_subscription**
> Subscription find_subscription(external_id)

Retrieve a subscription

This endpoint retrieves a specific subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.subscription import Subscription
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
    api_instance = lago_api.SubscriptionsApi(api_client)
    external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing subscription

    try:
        # Retrieve a subscription
        api_response = api_instance.find_subscription(external_id)
        print("The response of SubscriptionsApi->find_subscription:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubscriptionsApi->find_subscription: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_id** | **str**| External ID of the existing subscription | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Subscription |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_subscription**
> Subscription update_subscription(external_id, subscription_update_input)

Update a subscription

This endpoint allows you to update a subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.subscription import Subscription
from lago_api.models.subscription_update_input import SubscriptionUpdateInput
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
    api_instance = lago_api.SubscriptionsApi(api_client)
    external_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External ID of the existing subscription
    subscription_update_input = lago_api.SubscriptionUpdateInput() # SubscriptionUpdateInput | Update an existing subscription

    try:
        # Update a subscription
        api_response = api_instance.update_subscription(external_id, subscription_update_input)
        print("The response of SubscriptionsApi->update_subscription:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SubscriptionsApi->update_subscription: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_id** | **str**| External ID of the existing subscription | 
 **subscription_update_input** | [**SubscriptionUpdateInput**](SubscriptionUpdateInput.md)| Update an existing subscription | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Subscription updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

