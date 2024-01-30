# OpenAPI\Client\EventsApi

All URIs are relative to https://api.getlago.com/api/v1, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**createBatchEvents()**](EventsApi.md#createBatchEvents) | **POST** /events/batch | Batch multiple events |
| [**createEvent()**](EventsApi.md#createEvent) | **POST** /events | Send usage events |
| [**eventEstimateFees()**](EventsApi.md#eventEstimateFees) | **POST** /events/estimate_fees | Estimate fees for a pay in advance charge |
| [**findEvent()**](EventsApi.md#findEvent) | **GET** /events/{transaction_id} | Retrieve a specific event |


## `createBatchEvents()`

```php
createBatchEvents($event_batch_input)
```

Batch multiple events

This endpoint is used for transmitting a batch of usage measurement.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\EventsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$event_batch_input = new \OpenAPI\Client\Model\EventBatchInput(); // \OpenAPI\Client\Model\EventBatchInput | Batch events payload

try {
    $apiInstance->createBatchEvents($event_batch_input);
} catch (Exception $e) {
    echo 'Exception when calling EventsApi->createBatchEvents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **event_batch_input** | [**\OpenAPI\Client\Model\EventBatchInput**](../Model/EventBatchInput.md)| Batch events payload | |

### Return type

void (empty response body)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createEvent()`

```php
createEvent($event_input): \OpenAPI\Client\Model\Event
```

Send usage events

This endpoint is used for transmitting usage measurement events to either a designated customer or a specific subscription.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\EventsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$event_input = new \OpenAPI\Client\Model\EventInput(); // \OpenAPI\Client\Model\EventInput | Event payload

try {
    $result = $apiInstance->createEvent($event_input);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling EventsApi->createEvent: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **event_input** | [**\OpenAPI\Client\Model\EventInput**](../Model/EventInput.md)| Event payload | |

### Return type

[**\OpenAPI\Client\Model\Event**](../Model/Event.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `eventEstimateFees()`

```php
eventEstimateFees($event_estimate_fees_input): \OpenAPI\Client\Model\Fees
```

Estimate fees for a pay in advance charge

Estimate the fees that would be created after reception of an event for a billable metric attached to one or multiple pay in advance charges

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\EventsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$event_estimate_fees_input = new \OpenAPI\Client\Model\EventEstimateFeesInput(); // \OpenAPI\Client\Model\EventEstimateFeesInput | Event estimate payload

try {
    $result = $apiInstance->eventEstimateFees($event_estimate_fees_input);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling EventsApi->eventEstimateFees: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **event_estimate_fees_input** | [**\OpenAPI\Client\Model\EventEstimateFeesInput**](../Model/EventEstimateFeesInput.md)| Event estimate payload | |

### Return type

[**\OpenAPI\Client\Model\Fees**](../Model/Fees.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `findEvent()`

```php
findEvent($transaction_id): \OpenAPI\Client\Model\Event
```

Retrieve a specific event

This endpoint is used for retrieving a specific usage measurement event that has been sent to a customer or a subscription.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\EventsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$transaction_id = transaction_1234567890; // string | This field represents the unique identifier sent for this specific event.

try {
    $result = $apiInstance->findEvent($transaction_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling EventsApi->findEvent: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **transaction_id** | **string**| This field represents the unique identifier sent for this specific event. | |

### Return type

[**\OpenAPI\Client\Model\Event**](../Model/Event.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
