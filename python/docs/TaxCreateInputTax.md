# TaxCreateInputTax


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the tax. | 
**code** | **str** | Unique code used to identify the tax associated with the API request. | 
**rate** | **str** | The percentage rate of the tax | 
**description** | **str** | Internal description of the taxe | [optional] 
**applied_to_organization** | **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | [optional] 

## Example

```python
from lago_api.models.tax_create_input_tax import TaxCreateInputTax

# TODO update the JSON string below
json = "{}"
# create an instance of TaxCreateInputTax from a JSON string
tax_create_input_tax_instance = TaxCreateInputTax.from_json(json)
# print the JSON string representation of the object
print TaxCreateInputTax.to_json()

# convert the object into a dict
tax_create_input_tax_dict = tax_create_input_tax_instance.to_dict()
# create an instance of TaxCreateInputTax from a dict
tax_create_input_tax_form_dict = tax_create_input_tax.from_dict(tax_create_input_tax_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


