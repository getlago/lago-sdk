# lago_api.FeesApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**find_all_fees**](FeesApi.md#find_all_fees) | **GET** /fees | List all fees
[**find_fee**](FeesApi.md#find_fee) | **GET** /fees/{lago_id} | Retrieve a specific fee
[**update_fee**](FeesApi.md#update_fee) | **PUT** /fees/{lago_id} | Update a fee


# **find_all_fees**
> FeesPaginated find_all_fees(page=page, per_page=per_page, external_customer_id=external_customer_id, external_subscription_id=external_subscription_id, currency=currency, fee_type=fee_type, billable_metric_code=billable_metric_code, payment_status=payment_status, created_at_from=created_at_from, created_at_to=created_at_to, succeeded_at_from=succeeded_at_from, succeeded_at_to=succeeded_at_to, failed_at_from=failed_at_from, failed_at_to=failed_at_to, refunded_at_from=refunded_at_from, refunded_at_to=refunded_at_to)

List all fees

This endpoint is used for retrieving all fees that has been issued.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.fees_paginated import FeesPaginated
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
    api_instance = lago_api.FeesApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | Unique identifier assigned to the customer in your application. (optional)
    external_subscription_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | External subscription ID (optional)
    currency = lago_api.Currency() # Currency | Filter results by fee’s currency. (optional)
    fee_type = 'charge' # str | The fee type. Possible values are `add-on`, `charge`, `credit` or `subscription`. (optional)
    billable_metric_code = 'bm_code' # str | Filter results by the `code` of the billable metric attached to the fee. Only applies to `charge` types. (optional)
    payment_status = 'succeeded' # str | Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are `pending`, `succeeded`, `failed` and refunded`. (optional)
    created_at_from = '2023-03-28T12:21:51Z' # datetime | Filter results created after creation date and time in UTC. (optional)
    created_at_to = '2023-03-28T12:21:51Z' # datetime | Filter results created before creation date and time in UTC. (optional)
    succeeded_at_from = '2023-03-28T12:21:51Z' # datetime | Filter results with payment success after creation date and time in UTC. (optional)
    succeeded_at_to = '2023-03-28T12:21:51Z' # datetime | Filter results with payment success after creation date and time in UTC. (optional)
    failed_at_from = '2023-03-28T12:21:51Z' # datetime | Filter results with payment failure after creation date and time in UTC. (optional)
    failed_at_to = '2023-03-28T12:21:51Z' # datetime | Filter results with payment failure after creation date and time in UTC. (optional)
    refunded_at_from = '2023-03-28T12:21:51Z' # datetime | Filter results with payment refund after creation date and time in UTC. (optional)
    refunded_at_to = '2023-03-28T12:21:51Z' # datetime | Filter results with payment refund after creation date and time in UTC. (optional)

    try:
        # List all fees
        api_response = api_instance.find_all_fees(page=page, per_page=per_page, external_customer_id=external_customer_id, external_subscription_id=external_subscription_id, currency=currency, fee_type=fee_type, billable_metric_code=billable_metric_code, payment_status=payment_status, created_at_from=created_at_from, created_at_to=created_at_to, succeeded_at_from=succeeded_at_from, succeeded_at_to=succeeded_at_to, failed_at_from=failed_at_from, failed_at_to=failed_at_to, refunded_at_from=refunded_at_from, refunded_at_to=refunded_at_to)
        print("The response of FeesApi->find_all_fees:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FeesApi->find_all_fees: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **external_customer_id** | **str**| Unique identifier assigned to the customer in your application. | [optional] 
 **external_subscription_id** | **str**| External subscription ID | [optional] 
 **currency** | [**Currency**](.md)| Filter results by fee’s currency. | [optional] 
 **fee_type** | **str**| The fee type. Possible values are &#x60;add-on&#x60;, &#x60;charge&#x60;, &#x60;credit&#x60; or &#x60;subscription&#x60;. | [optional] 
 **billable_metric_code** | **str**| Filter results by the &#x60;code&#x60; of the billable metric attached to the fee. Only applies to &#x60;charge&#x60; types. | [optional] 
 **payment_status** | **str**| Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; and refunded&#x60;. | [optional] 
 **created_at_from** | **datetime**| Filter results created after creation date and time in UTC. | [optional] 
 **created_at_to** | **datetime**| Filter results created before creation date and time in UTC. | [optional] 
 **succeeded_at_from** | **datetime**| Filter results with payment success after creation date and time in UTC. | [optional] 
 **succeeded_at_to** | **datetime**| Filter results with payment success after creation date and time in UTC. | [optional] 
 **failed_at_from** | **datetime**| Filter results with payment failure after creation date and time in UTC. | [optional] 
 **failed_at_to** | **datetime**| Filter results with payment failure after creation date and time in UTC. | [optional] 
 **refunded_at_from** | **datetime**| Filter results with payment refund after creation date and time in UTC. | [optional] 
 **refunded_at_to** | **datetime**| Filter results with payment refund after creation date and time in UTC. | [optional] 

### Return type

[**FeesPaginated**](FeesPaginated.md)

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
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_fee**
> Fee find_fee(lago_id)

Retrieve a specific fee

This endpoint is used for retrieving a specific fee that has been issued.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.fee import Fee
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
    api_instance = lago_api.FeesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.

    try:
        # Retrieve a specific fee
        api_response = api_instance.find_fee(lago_id)
        print("The response of FeesApi->find_fee:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FeesApi->find_fee: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 

### Return type

[**Fee**](Fee.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Fee |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_fee**
> Fee update_fee(lago_id, fee_update_input=fee_update_input)

Update a fee

This endpoint is used for updating a specific fee that has been issued.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.fee import Fee
from lago_api.models.fee_update_input import FeeUpdateInput
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
    api_instance = lago_api.FeesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.
    fee_update_input = lago_api.FeeUpdateInput() # FeeUpdateInput | Fee payload (optional)

    try:
        # Update a fee
        api_response = api_instance.update_fee(lago_id, fee_update_input=fee_update_input)
        print("The response of FeesApi->update_fee:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FeesApi->update_fee: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system. | 
 **fee_update_input** | [**FeeUpdateInput**](FeeUpdateInput.md)| Fee payload | [optional] 

### Return type

[**Fee**](Fee.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Fee updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

