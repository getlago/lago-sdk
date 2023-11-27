# GroupPropertiesObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_id** | **str** | Unique identifier of a billable metric group, created by Lago. | 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual group will be used as the default display name. | [optional] 
**values** | [**GroupPropertiesObjectValues**](GroupPropertiesObjectValues.md) |  | 

## Example

```python
from lago_api.models.group_properties_object import GroupPropertiesObject

# TODO update the JSON string below
json = "{}"
# create an instance of GroupPropertiesObject from a JSON string
group_properties_object_instance = GroupPropertiesObject.from_json(json)
# print the JSON string representation of the object
print GroupPropertiesObject.to_json()

# convert the object into a dict
group_properties_object_dict = group_properties_object_instance.to_dict()
# create an instance of GroupPropertiesObject from a dict
group_properties_object_form_dict = group_properties_object.from_dict(group_properties_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


