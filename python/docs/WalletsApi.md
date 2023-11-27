# lago_api.WalletsApi

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_wallet**](WalletsApi.md#create_wallet) | **POST** /wallets | Create a wallet
[**create_wallet_transaction**](WalletsApi.md#create_wallet_transaction) | **POST** /wallet_transactions | Top up a wallet
[**destroy_wallet**](WalletsApi.md#destroy_wallet) | **DELETE** /wallets/{lago_id} | Terminate a wallet
[**find_all_wallet_transactions**](WalletsApi.md#find_all_wallet_transactions) | **GET** /wallets/{lago_id}/wallet_transactions | List all wallet transactions
[**find_all_wallets**](WalletsApi.md#find_all_wallets) | **GET** /wallets | List all wallets
[**find_wallet**](WalletsApi.md#find_wallet) | **GET** /wallets/{lago_id} | Retrieve a wallet
[**update_wallet**](WalletsApi.md#update_wallet) | **PUT** /wallets/{lago_id} | Update a wallet


# **create_wallet**
> Wallet create_wallet(wallet_create_input)

Create a wallet

This endpoint is used to create a wallet with prepaid credits.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet import Wallet
from lago_api.models.wallet_create_input import WalletCreateInput
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
    api_instance = lago_api.WalletsApi(api_client)
    wallet_create_input = lago_api.WalletCreateInput() # WalletCreateInput | Wallet payload

    try:
        # Create a wallet
        api_response = api_instance.create_wallet(wallet_create_input)
        print("The response of WalletsApi->create_wallet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->create_wallet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet_create_input** | [**WalletCreateInput**](WalletCreateInput.md)| Wallet payload | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_wallet_transaction**
> WalletTransactions create_wallet_transaction(wallet_transaction_create_input)

Top up a wallet

This endpoint is used to top-up an active wallet.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet_transaction_create_input import WalletTransactionCreateInput
from lago_api.models.wallet_transactions import WalletTransactions
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
    api_instance = lago_api.WalletsApi(api_client)
    wallet_transaction_create_input = lago_api.WalletTransactionCreateInput() # WalletTransactionCreateInput | Wallet transaction payload

    try:
        # Top up a wallet
        api_response = api_instance.create_wallet_transaction(wallet_transaction_create_input)
        print("The response of WalletsApi->create_wallet_transaction:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->create_wallet_transaction: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet_transaction_create_input** | [**WalletTransactionCreateInput**](WalletTransactionCreateInput.md)| Wallet transaction payload | 

### Return type

[**WalletTransactions**](WalletTransactions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet transaction created |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **destroy_wallet**
> Wallet destroy_wallet(lago_id)

Terminate a wallet

This endpoint is used to terminate an existing wallet with prepaid credits.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet import Wallet
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
    api_instance = lago_api.WalletsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.

    try:
        # Terminate a wallet
        api_response = api_instance.destroy_wallet(lago_id)
        print("The response of WalletsApi->destroy_wallet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->destroy_wallet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet terminated |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**405** | Not Allowed error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_wallet_transactions**
> WalletTransactionsPaginated find_all_wallet_transactions(lago_id, page=page, per_page=per_page, status=status, transaction_type=transaction_type)

List all wallet transactions

This endpoint is used to list all wallet transactions.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet_transactions_paginated import WalletTransactionsPaginated
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
    api_instance = lago_api.WalletsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)
    status = 'pending' # str | The status of the wallet transaction. Possible values are `pending` or `settled`. (optional)
    transaction_type = 'inbound' # str | The transaction type of the wallet transaction. Possible values are `inbound` (increasing the wallet balance) or `outbound` (decreasing the wallet balance). (optional)

    try:
        # List all wallet transactions
        api_response = api_instance.find_all_wallet_transactions(lago_id, page=page, per_page=per_page, status=status, transaction_type=transaction_type)
        print("The response of WalletsApi->find_all_wallet_transactions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->find_all_wallet_transactions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 
 **status** | **str**| The status of the wallet transaction. Possible values are &#x60;pending&#x60; or &#x60;settled&#x60;. | [optional] 
 **transaction_type** | **str**| The transaction type of the wallet transaction. Possible values are &#x60;inbound&#x60; (increasing the wallet balance) or &#x60;outbound&#x60; (decreasing the wallet balance). | [optional] 

### Return type

[**WalletTransactionsPaginated**](WalletTransactionsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet transactions |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_wallets**
> WalletsPaginated find_all_wallets(external_customer_id, page=page, per_page=per_page)

List all wallets

This endpoint is used to list all wallets with prepaid credits.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallets_paginated import WalletsPaginated
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
    api_instance = lago_api.WalletsApi(api_client)
    external_customer_id = '5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba' # str | The customer external unique identifier (provided by your own application).
    page = 1 # int | Page number. (optional)
    per_page = 20 # int | Number of records per page. (optional)

    try:
        # List all wallets
        api_response = api_instance.find_all_wallets(external_customer_id, page=page, per_page=per_page)
        print("The response of WalletsApi->find_all_wallets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->find_all_wallets: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **external_customer_id** | **str**| The customer external unique identifier (provided by your own application). | 
 **page** | **int**| Page number. | [optional] 
 **per_page** | **int**| Number of records per page. | [optional] 

### Return type

[**WalletsPaginated**](WalletsPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallets |  -  |
**401** | Unauthorized error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_wallet**
> Wallet find_wallet(lago_id)

Retrieve a wallet

This endpoint is used to retrieve an existing wallet with prepaid credits.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet import Wallet
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
    api_instance = lago_api.WalletsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.

    try:
        # Retrieve a wallet
        api_response = api_instance.find_wallet(lago_id)
        print("The response of WalletsApi->find_wallet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->find_wallet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_wallet**
> Wallet update_wallet(lago_id, wallet_update_input)

Update a wallet

This endpoint is used to update an existing wallet with prepaid credits.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import os
import lago_api
from lago_api.models.wallet import Wallet
from lago_api.models.wallet_update_input import WalletUpdateInput
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
    api_instance = lago_api.WalletsApi(api_client)
    lago_id = '1a901a90-1a90-1a90-1a90-1a901a901a90' # str | Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system.
    wallet_update_input = lago_api.WalletUpdateInput() # WalletUpdateInput | Wallet update payload

    try:
        # Update a wallet
        api_response = api_instance.update_wallet(lago_id, wallet_update_input)
        print("The response of WalletsApi->update_wallet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WalletsApi->update_wallet: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lago_id** | **str**| Unique identifier assigned to the wallet within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the wallet’s record within the Lago system. | 
 **wallet_update_input** | [**WalletUpdateInput**](WalletUpdateInput.md)| Wallet update payload | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Wallet updated |  -  |
**400** | Bad Request error |  -  |
**401** | Unauthorized error |  -  |
**404** | Not Found error |  -  |
**422** | Unprocessable entity error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

