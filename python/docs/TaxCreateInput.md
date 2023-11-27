# TaxCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tax** | [**TaxCreateInputTax**](TaxCreateInputTax.md) |  | 

## Example

```python
from lago_api.models.tax_create_input import TaxCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of TaxCreateInput from a JSON string
tax_create_input_instance = TaxCreateInput.from_json(json)
# print the JSON string representation of the object
print TaxCreateInput.to_json()

# convert the object into a dict
tax_create_input_dict = tax_create_input_instance.to_dict()
# create an instance of TaxCreateInput from a dict
tax_create_input_form_dict = tax_create_input.from_dict(tax_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


