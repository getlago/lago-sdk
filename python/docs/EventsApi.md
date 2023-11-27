# lago_api.EventsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_batch_events**](EventsApi.md#create_batch_events) | **POST** /events/batch | Batch multiple events
[**create_event**](EventsApi.md#create_event) | **POST** /events | Send usage events
[**event_estimate_fees**](EventsApi.md#event_estimate_fees) | **POST** /events/estimate_fees | Estimate fees for a pay in advance charge
[**find_event**](EventsApi.md#find_event) | **GET** /events/{transaction_id} | Retrieve a specific event


# **create_batch_events**
> create_batch_events(event_batch_input)

Batch multiple events

This endpoint is used for transmitting a batch of usage measurement events to multiple subscriptions for a single customer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.event_batch_input import EventBatchInput
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
    api_instance = lago_api.EventsApi(api_client)
    event_batch_input = lago_api.EventBatchInput() # EventBatchInput | Batch events payload

    try:
        # Batch multiple events
        api_instance.create_batch_events(event_batch_input)
    except Exception as e:
        print("Exception when calling EventsApi->create_batch_events: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event_batch_input** | [**EventBatchInput**](EventBatchInput.md)| Batch events payload | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Event received |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_event**
> Event create_event(event_input)

Send usage events

This endpoint is used for transmitting usage measurement events to either a designated customer or a specific subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.event import Event
from lago_api.models.event_input import EventInput
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
    api_instance = lago_api.EventsApi(api_client)
    event_input = lago_api.EventInput() # EventInput | Event payload

    try:
        # Send usage events
        api_response = api_instance.create_event(event_input)
        print("The response of EventsApi->create_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EventsApi->create_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event_input** | [**EventInput**](EventInput.md)| Event payload | 

### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Event |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **event_estimate_fees**
> Fees event_estimate_fees(event_estimate_fees_input)

Estimate fees for a pay in advance charge

Estimate the fees that would be created after reception of an event for a billable metric attached to one or multiple pay in advance charges

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.event_estimate_fees_input import EventEstimateFeesInput
from lago_api.models.fees import Fees
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
    api_instance = lago_api.EventsApi(api_client)
    event_estimate_fees_input = lago_api.EventEstimateFeesInput() # EventEstimateFeesInput | Event estimate payload

    try:
        # Estimate fees for a pay in advance charge
        api_response = api_instance.event_estimate_fees(event_estimate_fees_input)
        print("The response of EventsApi->event_estimate_fees:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EventsApi->event_estimate_fees: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event_estimate_fees_input** | [**EventEstimateFeesInput**](EventEstimateFeesInput.md)| Event estimate payload | 

### Return type

[**Fees**](Fees.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Fees estimate |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_event**
> Event find_event(transaction_id)

Retrieve a specific event

This endpoint is used for retrieving a specific usage measurement event that has been sent to a customer or a subscription.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.event import Event
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
    api_instance = lago_api.EventsApi(api_client)
    transaction_id = 'transaction_1234567890' # str | This field represents the unique identifier sent for this specific event.

    try:
        # Retrieve a specific event
        api_response = api_instance.find_event(transaction_id)
        print("The response of EventsApi->find_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EventsApi->find_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction_id** | **str**| This field represents the unique identifier sent for this specific event. | 

### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Event |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

