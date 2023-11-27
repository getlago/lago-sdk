# EventEstimateFeesInputEvent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **str** | The code that identifies a targeted billable metric. It is essential that this code matches the &#x60;code&#x60; property of one of your active billable metrics. If the provided code does not correspond to any active billable metric, it will be ignored during the process. | 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application). This field is optional if you send the &#x60;external_subscription_id&#x60;, targeting a specific subscription. | [optional] 
**external_subscription_id** | **str** | The unique identifier of the subscription within your application. It is a mandatory field when the customer possesses multiple subscriptions or when the &#x60;external_customer_id&#x60; is not provided. | [optional] 
**properties** | **object** | This field represents additional properties associated with the event, which are utilized in the calculation of the final fee. This object becomes mandatory when the targeted billable metric employs a &#x60;sum_agg&#x60;, &#x60;max_agg&#x60;, or &#x60;unique_count_agg&#x60; aggregation method. However, when using a simple &#x60;count_agg&#x60;, this object is not required. | [optional] 

## Example

```python
from lago_api.models.event_estimate_fees_input_event import EventEstimateFeesInputEvent

# TODO update the JSON string below
json = "{}"
# create an instance of EventEstimateFeesInputEvent from a JSON string
event_estimate_fees_input_event_instance = EventEstimateFeesInputEvent.from_json(json)
# print the JSON string representation of the object
print EventEstimateFeesInputEvent.to_json()

# convert the object into a dict
event_estimate_fees_input_event_dict = event_estimate_fees_input_event_instance.to_dict()
# create an instance of EventEstimateFeesInputEvent from a dict
event_estimate_fees_input_event_form_dict = event_estimate_fees_input_event.from_dict(event_estimate_fees_input_event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


