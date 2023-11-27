# GroupObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of a specific group associated with the billable metric. | 
**key** | **str** | Key of a specific group associated with the billable metric. | 
**value** | **str** | One of the values for a specific group associated with the billable metric. | 

## Example

```python
from lago_api.models.group_object import GroupObject

# TODO update the JSON string below
json = "{}"
# create an instance of GroupObject from a JSON string
group_object_instance = GroupObject.from_json(json)
# print the JSON string representation of the object
print GroupObject.to_json()

# convert the object into a dict
group_object_dict = group_object_instance.to_dict()
# create an instance of GroupObject from a dict
group_object_form_dict = group_object.from_dict(group_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


