# AddOn


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**add_on** | [**AddOnObject**](AddOnObject.md) |  | 

## Example

```python
from lago_api.models.add_on import AddOn

# TODO update the JSON string below
json = "{}"
# create an instance of AddOn from a JSON string
add_on_instance = AddOn.from_json(json)
# print the JSON string representation of the object
print AddOn.to_json()

# convert the object into a dict
add_on_dict = add_on_instance.to_dict()
# create an instance of AddOn from a dict
add_on_form_dict = add_on.from_dict(add_on_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


