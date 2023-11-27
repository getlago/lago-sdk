# lago_api.PlansApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_plan**](PlansApi.md#create_plan) | **POST** /plans | Create a plan
[**destroy_plan**](PlansApi.md#destroy_plan) | **DELETE** /plans/{code} | Delete a plan
[**find_all_plans**](PlansApi.md#find_all_plans) | **GET** /plans | List all plans
[**find_plan**](PlansApi.md#find_plan) | **GET** /plans/{code} | Retrieve a plan
[**update_plan**](PlansApi.md#update_plan) | **PUT** /plans/{code} | Update a plan


# **create_plan**
> Plan create_plan(plan_create_input)

Create a plan

This endpoint creates a plan with subscription and usage-based charges. It supports flexible billing cadence (in-advance or in-arrears) and allows for both recurring and metered charges.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.plan import Plan
from lago_api.models.plan_create_input import PlanCreateInput
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
    api_instance = lago_api.PlansApi(api_client)
    plan_create_input = lago_api.PlanCreateInput() # PlanCreateInput | Plan payload

    try:
        # Create a plan
        api_response = api_instance.create_plan(plan_create_input)
        print("The response of PlansApi->create_plan:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlansApi->create_plan: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan_create_input** | [**PlanCreateInput**](PlanCreateInput.md)| Plan payload | 

### Return type

[**Plan**](Plan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Plan created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_plan**
> Plan destroy_plan(code)

Delete a plan

This endpoint deletes a specific plan. Note that this plan could be associated with active subscriptions.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.plan import Plan
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
    api_instance = lago_api.PlansApi(api_client)
    code = 'startup' # str | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance.

    try:
        # Delete a plan
        api_response = api_instance.destroy_plan(code)
        print("The response of PlansApi->destroy_plan:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlansApi->destroy_plan: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | 

### Return type

[**Plan**](Plan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Plan deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_plans**
> PlansPaginated find_all_plans(page=page, per_page=per_page)

List all plans

This endpoint retrieves all existing plans.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.plans_paginated import PlansPaginated
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
    api_instance = lago_api.PlansApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all plans
        api_response = api_instance.find_all_plans(page=page, per_page=per_page)
        print("The response of PlansApi->find_all_plans:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlansApi->find_all_plans: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**PlansPaginated**](PlansPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Plans |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_plan**
> Plan find_plan(code)

Retrieve a plan

This endpoint retrieves a specific plan.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.plan import Plan
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
    api_instance = lago_api.PlansApi(api_client)
    code = 'startup' # str | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance.

    try:
        # Retrieve a plan
        api_response = api_instance.find_plan(code)
        print("The response of PlansApi->find_plan:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlansApi->find_plan: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | 

### Return type

[**Plan**](Plan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Plan |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_plan**
> Plan update_plan(code, plan_update_input)

Update a plan

This endpoint updates a specific plan with subscription and usage-based charges. It supports flexible billing cadence (in-advance or in-arrears) and allows for both recurring and metered charges.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.plan import Plan
from lago_api.models.plan_update_input import PlanUpdateInput
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
    api_instance = lago_api.PlansApi(api_client)
    code = 'startup' # str | The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance.
    plan_update_input = lago_api.PlanUpdateInput() # PlanUpdateInput | Plan payload

    try:
        # Update a plan
        api_response = api_instance.update_plan(code, plan_update_input)
        print("The response of PlansApi->update_plan:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlansApi->update_plan: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| The code of the plan. It serves as a unique identifier associated with a particular plan. The code is typically used for internal or system-level identification purposes, like assigning a subscription, for instance. | 
 **plan_update_input** | [**PlanUpdateInput**](PlanUpdateInput.md)| Plan payload | 

### Return type

[**Plan**](Plan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Plan updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

