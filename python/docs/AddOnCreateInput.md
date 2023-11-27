# AddOnCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**add_on** | [**AddOnCreateInputAddOn**](AddOnCreateInputAddOn.md) |  | 

## Example

```python
from lago_api.models.add_on_create_input import AddOnCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AddOnCreateInput from a JSON string
add_on_create_input_instance = AddOnCreateInput.from_json(json)
# print the JSON string representation of the object
print AddOnCreateInput.to_json()

# convert the object into a dict
add_on_create_input_dict = add_on_create_input_instance.to_dict()
# create an instance of AddOnCreateInput from a dict
add_on_create_input_form_dict = add_on_create_input.from_dict(add_on_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


