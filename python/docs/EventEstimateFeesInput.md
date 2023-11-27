# EventEstimateFeesInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**event** | [**EventEstimateFeesInputEvent**](EventEstimateFeesInputEvent.md) |  | 

## Example

```python
from lago_api.models.event_estimate_fees_input import EventEstimateFeesInput

# TODO update the JSON string below
json = "{}"
# create an instance of EventEstimateFeesInput from a JSON string
event_estimate_fees_input_instance = EventEstimateFeesInput.from_json(json)
# print the JSON string representation of the object
print EventEstimateFeesInput.to_json()

# convert the object into a dict
event_estimate_fees_input_dict = event_estimate_fees_input_instance.to_dict()
# create an instance of EventEstimateFeesInput from a dict
event_estimate_fees_input_form_dict = event_estimate_fees_input.from_dict(event_estimate_fees_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


