# AddOnBaseInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the add-on. | [optional] 
**invoice_display_name** | **str** | Specifies the name that will be displayed on an invoice. If no value is set for this field, the name of the actual charge will be used as the default display name. | [optional] 
**code** | **str** | Unique code used to identify the add-on. | [optional] 
**amount_cents** | **int** | The cost of the add-on in cents, excluding any applicable taxes, that is billed to a customer. By creating a one-off invoice, you will be able to override this value. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**description** | **str** | The description of the add-on. | [optional] 
**tax_codes** | **List[str]** | List of unique code used to identify the taxes. | [optional] 

## Example

```python
from lago_api.models.add_on_base_input import AddOnBaseInput

# TODO update the JSON string below
json = "{}"
# create an instance of AddOnBaseInput from a JSON string
add_on_base_input_instance = AddOnBaseInput.from_json(json)
# print the JSON string representation of the object
print AddOnBaseInput.to_json()

# convert the object into a dict
add_on_base_input_dict = add_on_base_input_instance.to_dict()
# create an instance of AddOnBaseInput from a dict
add_on_base_input_form_dict = add_on_base_input.from_dict(add_on_base_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


