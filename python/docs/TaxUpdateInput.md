# TaxUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tax** | [**TaxBaseInput**](TaxBaseInput.md) |  | 

## Example

```python
from lago_api.models.tax_update_input import TaxUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of TaxUpdateInput from a JSON string
tax_update_input_instance = TaxUpdateInput.from_json(json)
# print the JSON string representation of the object
print TaxUpdateInput.to_json()

# convert the object into a dict
tax_update_input_dict = tax_update_input_instance.to_dict()
# create an instance of TaxUpdateInput from a dict
tax_update_input_form_dict = tax_update_input.from_dict(tax_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


