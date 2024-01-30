# OpenAPI\Client\SubscriptionsApi

All URIs are relative to https://api.getlago.com/api/v1, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**createSubscription()**](SubscriptionsApi.md#createSubscription) | **POST** /subscriptions | Assign a plan to a customer |
| [**destroySubscription()**](SubscriptionsApi.md#destroySubscription) | **DELETE** /subscriptions/{external_id} | Terminate a subscription |
| [**findAllSubscriptions()**](SubscriptionsApi.md#findAllSubscriptions) | **GET** /subscriptions | List all subscriptions |
| [**findSubscription()**](SubscriptionsApi.md#findSubscription) | **GET** /subscriptions/{external_id} | Retrieve a subscription |
| [**updateSubscription()**](SubscriptionsApi.md#updateSubscription) | **PUT** /subscriptions/{external_id} | Update a subscription |


## `createSubscription()`

```php
createSubscription($subscription_create_input): \OpenAPI\Client\Model\Subscription
```

Assign a plan to a customer

This endpoint assigns a plan to a customer, creating or modifying a subscription. Ideal for initial subscriptions or plan changes (upgrades/downgrades).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SubscriptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$subscription_create_input = new \OpenAPI\Client\Model\SubscriptionCreateInput(); // \OpenAPI\Client\Model\SubscriptionCreateInput | Subscription payload

try {
    $result = $apiInstance->createSubscription($subscription_create_input);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionsApi->createSubscription: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **subscription_create_input** | [**\OpenAPI\Client\Model\SubscriptionCreateInput**](../Model/SubscriptionCreateInput.md)| Subscription payload | |

### Return type

[**\OpenAPI\Client\Model\Subscription**](../Model/Subscription.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `destroySubscription()`

```php
destroySubscription($external_id, $status): \OpenAPI\Client\Model\Subscription
```

Terminate a subscription

This endpoint allows you to terminate a subscription.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SubscriptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$external_id = 5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba; // string | External ID of the existing subscription
$status = pending; // string | If the field is not defined, Lago will terminate only `active` subscriptions. However, if you wish to cancel a `pending` subscription, please ensure that you include `status=pending` in your request.

try {
    $result = $apiInstance->destroySubscription($external_id, $status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionsApi->destroySubscription: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **external_id** | **string**| External ID of the existing subscription | |
| **status** | **string**| If the field is not defined, Lago will terminate only &#x60;active&#x60; subscriptions. However, if you wish to cancel a &#x60;pending&#x60; subscription, please ensure that you include &#x60;status&#x3D;pending&#x60; in your request. | [optional] |

### Return type

[**\OpenAPI\Client\Model\Subscription**](../Model/Subscription.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `findAllSubscriptions()`

```php
findAllSubscriptions($page, $per_page, $external_customer_id, $plan_code, $status): \OpenAPI\Client\Model\SubscriptionsPaginated
```

List all subscriptions

This endpoint retrieves all active subscriptions.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SubscriptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$page = 1; // int | Page number.
$per_page = 20; // int | Number of records per page.
$external_customer_id = 5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba; // string | The customer external unique identifier (provided by your own application)
$plan_code = premium; // string | The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans.
$status = ["active","pending"]; // string[] | If the field is not defined, Lago will return only `active` subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: `pending`, `canceled`, `terminated`, `active`.

try {
    $result = $apiInstance->findAllSubscriptions($page, $per_page, $external_customer_id, $plan_code, $status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionsApi->findAllSubscriptions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **page** | **int**| Page number. | [optional] |
| **per_page** | **int**| Number of records per page. | [optional] |
| **external_customer_id** | **string**| The customer external unique identifier (provided by your own application) | [optional] |
| **plan_code** | **string**| The unique code representing the plan to be attached to the customer. This code must correspond to the code property of one of the active plans. | [optional] |
| **status** | [**string[]**](../Model/string.md)| If the field is not defined, Lago will return only &#x60;active&#x60; subscriptions. However, if you wish to fetch subscriptions by different status you can define them in a status[] query param. Available filter values: &#x60;pending&#x60;, &#x60;canceled&#x60;, &#x60;terminated&#x60;, &#x60;active&#x60;. | [optional] |

### Return type

[**\OpenAPI\Client\Model\SubscriptionsPaginated**](../Model/SubscriptionsPaginated.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `findSubscription()`

```php
findSubscription($external_id): \OpenAPI\Client\Model\Subscription
```

Retrieve a subscription

This endpoint retrieves a specific subscription.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SubscriptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$external_id = 5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba; // string | External ID of the existing subscription

try {
    $result = $apiInstance->findSubscription($external_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionsApi->findSubscription: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **external_id** | **string**| External ID of the existing subscription | |

### Return type

[**\OpenAPI\Client\Model\Subscription**](../Model/Subscription.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateSubscription()`

```php
updateSubscription($external_id, $subscription_update_input): \OpenAPI\Client\Model\Subscription
```

Update a subscription

This endpoint allows you to update a subscription.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SubscriptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$external_id = 5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba; // string | External ID of the existing subscription
$subscription_update_input = new \OpenAPI\Client\Model\SubscriptionUpdateInput(); // \OpenAPI\Client\Model\SubscriptionUpdateInput | Update an existing subscription

try {
    $result = $apiInstance->updateSubscription($external_id, $subscription_update_input);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubscriptionsApi->updateSubscription: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **external_id** | **string**| External ID of the existing subscription | |
| **subscription_update_input** | [**\OpenAPI\Client\Model\SubscriptionUpdateInput**](../Model/SubscriptionUpdateInput.md)| Update an existing subscription | |

### Return type

[**\OpenAPI\Client\Model\Subscription**](../Model/Subscription.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
