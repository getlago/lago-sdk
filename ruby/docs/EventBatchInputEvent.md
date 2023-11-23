# LagoAPI::EventBatchInputEvent

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **transaction_id** | **String** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. |  |
| **external_customer_id** | **String** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_ids&#x60;, targeting a specific subscription. | [optional] |
| **external_subscription_ids** | **Array&lt;String&gt;** | Array of unique identifiers of the targeted subscriptions within your application. |  |
| **code** | **String** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. |  |
| **timestamp** | **Integer** | This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. | [optional] |
| **properties** | [**EventBatchInputEventProperties**](EventBatchInputEventProperties.md) |  | [optional] |

## Example

```ruby
require 'lago_ruby'

instance = LagoAPI::EventBatchInputEvent.new(
  transaction_id: transaction_1234567890,
  external_customer_id: 5eb02857-a71e-4ea2-bcf9-57d3a41bc6ba,
  external_subscription_ids: [&quot;sub_1234567890&quot;,&quot;sub_0987654321&quot;],
  code: storage,
  timestamp: 1651240791,
  properties: null
)
```

