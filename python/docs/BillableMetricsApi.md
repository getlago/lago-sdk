# lago_api.BillableMetricsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_billable_metric**](BillableMetricsApi.md#create_billable_metric) | **POST** /billable_metrics | Create a billable metric
[**destroy_billable_metric**](BillableMetricsApi.md#destroy_billable_metric) | **DELETE** /billable_metrics/{code} | Delete a billable metric
[**find_all_billable_metric_groups**](BillableMetricsApi.md#find_all_billable_metric_groups) | **GET** /billable_metrics/{code}/groups | Find a billable metric&#39;s groups
[**find_all_billable_metrics**](BillableMetricsApi.md#find_all_billable_metrics) | **GET** /billable_metrics | List all billable metrics
[**find_billable_metric**](BillableMetricsApi.md#find_billable_metric) | **GET** /billable_metrics/{code} | Retrieve a billable metric
[**update_billable_metric**](BillableMetricsApi.md#update_billable_metric) | **PUT** /billable_metrics/{code} | Update a billable metric


# **create_billable_metric**
> BillableMetric create_billable_metric(billable_metric_create_input)

Create a billable metric

This endpoint creates a new billable metric representing a pricing component of your application.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.billable_metric import BillableMetric
from lago_api.models.billable_metric_create_input import BillableMetricCreateInput
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    billable_metric_create_input = lago_api.BillableMetricCreateInput() # BillableMetricCreateInput | Billable metric payload

    try:
        # Create a billable metric
        api_response = api_instance.create_billable_metric(billable_metric_create_input)
        print("The response of BillableMetricsApi->create_billable_metric:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->create_billable_metric: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billable_metric_create_input** | [**BillableMetricCreateInput**](BillableMetricCreateInput.md)| Billable metric payload | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Billable metric created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_billable_metric**
> BillableMetric destroy_billable_metric(code)

Delete a billable metric

This endpoint deletes an existing billable metric representing a pricing component of your application.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.billable_metric import BillableMetric
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    code = 'storage' # str | Code of the existing billable metric.

    try:
        # Delete a billable metric
        api_response = api_instance.destroy_billable_metric(code)
        print("The response of BillableMetricsApi->destroy_billable_metric:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->destroy_billable_metric: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Code of the existing billable metric. | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Billable metric deleted |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_billable_metric_groups**
> GroupsPaginated find_all_billable_metric_groups(code, page=page, per_page=per_page)

Find a billable metric's groups

This endpoint retrieves all groups for a billable metric.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.groups_paginated import GroupsPaginated
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    code = 'example_code' # str | Code of the existing billable metric.
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # Find a billable metric's groups
        api_response = api_instance.find_all_billable_metric_groups(code, page=page, per_page=per_page)
        print("The response of BillableMetricsApi->find_all_billable_metric_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->find_all_billable_metric_groups: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Code of the existing billable metric. | 
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**GroupsPaginated**](GroupsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of billable metric&#39;s groups |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_billable_metrics**
> BillableMetricsPaginated find_all_billable_metrics(page=page, per_page=per_page)

List all billable metrics

This endpoint retrieves all existing billable metrics that represent pricing components of your application.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.billable_metrics_paginated import BillableMetricsPaginated
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all billable metrics
        api_response = api_instance.find_all_billable_metrics(page=page, per_page=per_page)
        print("The response of BillableMetricsApi->find_all_billable_metrics:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->find_all_billable_metrics: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**BillableMetricsPaginated**](BillableMetricsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of billable metrics |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_billable_metric**
> BillableMetric find_billable_metric(code)

Retrieve a billable metric

This endpoint retrieves an existing billable metric that represents a pricing component of your application. The billable metric is identified by its unique code.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.billable_metric import BillableMetric
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    code = 'storage' # str | Code of the existing billable metric.

    try:
        # Retrieve a billable metric
        api_response = api_instance.find_billable_metric(code)
        print("The response of BillableMetricsApi->find_billable_metric:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->find_billable_metric: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Code of the existing billable metric. | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Billable metric |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_billable_metric**
> BillableMetric update_billable_metric(code, billable_metric_update_input)

Update a billable metric

This endpoint updates an existing billable metric representing a pricing component of your application.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.billable_metric import BillableMetric
from lago_api.models.billable_metric_update_input import BillableMetricUpdateInput
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
    api_instance = lago_api.BillableMetricsApi(api_client)
    code = 'storage' # str | Code of the existing billable metric.
    billable_metric_update_input = lago_api.BillableMetricUpdateInput() # BillableMetricUpdateInput | Billable metric payload

    try:
        # Update a billable metric
        api_response = api_instance.update_billable_metric(code, billable_metric_update_input)
        print("The response of BillableMetricsApi->update_billable_metric:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling BillableMetricsApi->update_billable_metric: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **str**| Code of the existing billable metric. | 
 **billable_metric_update_input** | [**BillableMetricUpdateInput**](BillableMetricUpdateInput.md)| Billable metric payload | 

### Return type

[**BillableMetric**](BillableMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Billable metric updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

