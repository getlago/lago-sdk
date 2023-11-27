# EventBatchInputEventProperties

This field represents additional properties associated with the event, which are utilized in the calculation of the final fee. This object becomes mandatory when the targeted billable metric employs a `sum_agg`, `max_agg`, or `unique_count_agg` aggregation method. However, when using a simple `count_agg`, this object is not required.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**operation_type** | **str** | The &#x60;operation_type&#x60; field is only necessary when adding or removing a specific unit when the targeted billable metric adopts a &#x60;unique_count_agg&#x60; aggregation method. In other cases, the &#x60;operation_type&#x60; field is not required. The valid values for the &#x60;operation_type&#x60; field are &#x60;add&#x60; or &#x60;remove&#x60;, which indicate whether the unit is being added or removed from the unique count aggregation, respectively. | [optional] 

## Example

```python
from lago_api.models.event_batch_input_event_properties import EventBatchInputEventProperties

# TODO update the JSON string below
json = "{}"
# create an instance of EventBatchInputEventProperties from a JSON string
event_batch_input_event_properties_instance = EventBatchInputEventProperties.from_json(json)
# print the JSON string representation of the object
print EventBatchInputEventProperties.to_json()

# convert the object into a dict
event_batch_input_event_properties_dict = event_batch_input_event_properties_instance.to_dict()
# create an instance of EventBatchInputEventProperties from a dict
event_batch_input_event_properties_form_dict = event_batch_input_event_properties.from_dict(event_batch_input_event_properties_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


