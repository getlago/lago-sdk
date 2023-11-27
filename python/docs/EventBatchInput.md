# EventBatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**event** | [**EventBatchInputEvent**](EventBatchInputEvent.md) |  | 

## Example

```python
from lago_api.models.event_batch_input import EventBatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of EventBatchInput from a JSON string
event_batch_input_instance = EventBatchInput.from_json(json)
# print the JSON string representation of the object
print EventBatchInput.to_json()

# convert the object into a dict
event_batch_input_dict = event_batch_input_instance.to_dict()
# create an instance of EventBatchInput from a dict
event_batch_input_form_dict = event_batch_input.from_dict(event_batch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


