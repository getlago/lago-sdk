# lago_api.CreditNotesApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_credit_note**](CreditNotesApi.md#create_credit_note) | **POST** /credit_notes | Create a credit note
[**download_credit_note**](CreditNotesApi.md#download_credit_note) | **POST** /credit_notes/{lago_id}/download | Download a credit note PDF
[**estimate_credit_note**](CreditNotesApi.md#estimate_credit_note) | **POST** /credit_notes/estimate | Estimate amounts for a new credit note
[**find_all_credit_notes**](CreditNotesApi.md#find_all_credit_notes) | **GET** /credit_notes | List all credit notes
[**find_credit_note**](CreditNotesApi.md#find_credit_note) | **GET** /credit_notes/{lago_id} | Retrieve a credit note
[**update_credit_note**](CreditNotesApi.md#update_credit_note) | **PUT** /credit_notes/{lago_id} | Update a credit note
[**void_credit_note**](CreditNotesApi.md#void_credit_note) | **PUT** /credit_notes/{lago_id}/void | Void a credit note


# **create_credit_note**
> CreditNote create_credit_note(credit_note_create_input)

Create a credit note

This endpoint creates a new credit note.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note import CreditNote
from lago_api.models.credit_note_create_input import CreditNoteCreateInput
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
    api_instance = lago_api.CreditNotesApi(api_client)
    credit_note_create_input = lago_api.CreditNoteCreateInput() # CreditNoteCreateInput | Credit note payload

    try:
        # Create a credit note
        api_response = api_instance.create_credit_note(credit_note_create_input)
        print("The response of CreditNotesApi->create_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->create_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credit_note_create_input** | [**CreditNoteCreateInput**](CreditNoteCreateInput.md)| Credit note payload | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_credit_note**
> CreditNote download_credit_note(lago_id)

Download a credit note PDF

This endpoint downloads the PDF of an existing credit note.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note import CreditNote
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
    api_instance = lago_api.CreditNotesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | The credit note unique identifier, created by Lago.

    try:
        # Download a credit note PDF
        api_response = api_instance.download_credit_note(lago_id)
        print("The response of CreditNotesApi->download_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->download_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| The credit note unique identifier, created by Lago. | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note PDF |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **estimate_credit_note**
> CreditNoteEstimated estimate_credit_note(credit_note_estimate_input=credit_note_estimate_input)

Estimate amounts for a new credit note

This endpoint allows you to retrieve amounts for a new credit note creation.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note_estimate_input import CreditNoteEstimateInput
from lago_api.models.credit_note_estimated import CreditNoteEstimated
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
    api_instance = lago_api.CreditNotesApi(api_client)
    credit_note_estimate_input = lago_api.CreditNoteEstimateInput() # CreditNoteEstimateInput | Credit note estimate payload (optional)

    try:
        # Estimate amounts for a new credit note
        api_response = api_instance.estimate_credit_note(credit_note_estimate_input=credit_note_estimate_input)
        print("The response of CreditNotesApi->estimate_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->estimate_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credit_note_estimate_input** | [**CreditNoteEstimateInput**](CreditNoteEstimateInput.md)| Credit note estimate payload | [optional] 

### Return type

[**CreditNoteEstimated**](CreditNoteEstimated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note amounts |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_credit_notes**
> CreditNotes find_all_credit_notes(page=page, per_page=per_page, external_customer_id=external_customer_id)

List all credit notes

This endpoint list all existing credit notes.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_notes import CreditNotes
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
    api_instance = lago_api.CreditNotesApi(api_client)
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | Unique identifier assigned to the customer in your application. (optional)

    try:
        # List all credit notes
        api_response = api_instance.find_all_credit_notes(page=page, per_page=per_page, external_customer_id=external_customer_id)
        print("The response of CreditNotesApi->find_all_credit_notes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->find_all_credit_notes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **external_customer_id** | **str**| Unique identifier assigned to the customer in your application. | [optional] 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit notes |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_credit_note**
> CreditNote find_credit_note(lago_id)

Retrieve a credit note

This endpoint retrieves an existing credit note.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note import CreditNote
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
    api_instance = lago_api.CreditNotesApi(api_client)
    lago_id = '12345' # str | The credit note unique identifier, created by Lago.

    try:
        # Retrieve a credit note
        api_response = api_instance.find_credit_note(lago_id)
        print("The response of CreditNotesApi->find_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->find_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| The credit note unique identifier, created by Lago. | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_credit_note**
> CreditNote update_credit_note(lago_id, credit_note_update_input)

Update a credit note

This endpoint updates an existing credit note.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note import CreditNote
from lago_api.models.credit_note_update_input import CreditNoteUpdateInput
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
    api_instance = lago_api.CreditNotesApi(api_client)
    lago_id = '12345' # str | The credit note unique identifier, created by Lago.
    credit_note_update_input = lago_api.CreditNoteUpdateInput() # CreditNoteUpdateInput | Credit note update payload

    try:
        # Update a credit note
        api_response = api_instance.update_credit_note(lago_id, credit_note_update_input)
        print("The response of CreditNotesApi->update_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->update_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| The credit note unique identifier, created by Lago. | 
 **credit_note_update_input** | [**CreditNoteUpdateInput**](CreditNoteUpdateInput.md)| Credit note update payload | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **void_credit_note**
> CreditNote void_credit_note(lago_id)

Void a credit note

This endpoint voids an existing credit note.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.credit_note import CreditNote
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
    api_instance = lago_api.CreditNotesApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | The credit note unique identifier, created by Lago.

    try:
        # Void a credit note
        api_response = api_instance.void_credit_note(lago_id)
        print("The response of CreditNotesApi->void_credit_note:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CreditNotesApi->void_credit_note: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| The credit note unique identifier, created by Lago. | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Credit note voided |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**405** | Not Allowed error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

