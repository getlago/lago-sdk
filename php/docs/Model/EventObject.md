# # EventObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **string** | Unique identifier assigned to the event within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the event&#39;s record within the Lago system |
**transaction_id** | **string** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. |
**lago_customer_id** | **string** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system |
**external_customer_id** | **string** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_id&#x60;, targeting a specific subscription. |
**code** | **string** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. |
**timestamp** | **\DateTime** | This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. |
**properties** | [**\OpenAPI\Client\Model\EventObjectProperties**](EventObjectProperties.md) |  | [optional]
**lago_subscription_id** | **string** | Unique identifier assigned to the subscription within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the subscription’s record within the Lago system |
**external_subscription_id** | **string** | The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the &#x60;external_customer_id&#x60; is not provided. |
**created_at** | **\DateTime** | The creation date of the event&#39;s record in the Lago application, presented in the ISO 8601 datetime format, specifically in Coordinated Universal Time (UTC). It provides the precise timestamp of when the event&#39;s record was created within the Lago application |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
