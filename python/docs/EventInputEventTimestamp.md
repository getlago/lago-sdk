# EventInputEventTimestamp

This field captures the Unix timestamp in seconds indicating the occurrence of the event in Coordinated Universal Time (UTC). If this timestamp is not provided, the API will automatically set it to the time of event reception. You can also provide miliseconds precision by appending decimals to the timestamp. 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

## Example

```python
from lago_api.models.event_input_event_timestamp import EventInputEventTimestamp

# TODO update the JSON string below
json = "{}"
# create an instance of EventInputEventTimestamp from a JSON string
event_input_event_timestamp_instance = EventInputEventTimestamp.from_json(json)
# print the JSON string representation of the object
print EventInputEventTimestamp.to_json()

# convert the object into a dict
event_input_event_timestamp_dict = event_input_event_timestamp_instance.to_dict()
# create an instance of EventInputEventTimestamp from a dict
event_input_event_timestamp_form_dict = event_input_event_timestamp.from_dict(event_input_event_timestamp_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


