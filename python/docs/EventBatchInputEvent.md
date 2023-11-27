# EventBatchInputEvent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**transaction_id** | **str** | This field represents a unique identifier for the event. It is crucial for ensuring idempotency, meaning that each event can be uniquely identified and processed without causing any unintended side effects. | 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_ids&#x60;, targeting a specific subscription. | [optional] 
**external_subscription_ids** | **List[str]** | Array of unique identifiers of the targeted subscriptions within your application. | 
**code** | **str** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. | 
**timestamp** | **int** | This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. | [optional] 
**properties** | [**EventBatchInputEventProperties**](EventBatchInputEventProperties.md) |  | [optional] 

## Example

```python
from lago_api.models.event_batch_input_event import EventBatchInputEvent

# TODO update the JSON string below
json = "{}"
# create an instance of EventBatchInputEvent from a JSON string
event_batch_input_event_instance = EventBatchInputEvent.from_json(json)
# print the JSON string representation of the object
print EventBatchInputEvent.to_json()

# convert the object into a dict
event_batch_input_event_dict = event_batch_input_event_instance.to_dict()
# create an instance of EventBatchInputEvent from a dict
event_batch_input_event_form_dict = event_batch_input_event.from_dict(event_batch_input_event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


