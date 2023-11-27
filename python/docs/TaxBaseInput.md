# TaxBaseInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the tax. | [optional] 
**code** | **str** | Unique code used to identify the tax associated with the API request. | [optional] 
**rate** | **str** | The percentage rate of the tax | [optional] 
**description** | **str** | Internal description of the taxe | [optional] 
**applied_to_organization** | **bool** | Set to &#x60;true&#x60; if the tax is used as one of the organization&#39;s default | [optional] 

## Example

```python
from lago_api.models.tax_base_input import TaxBaseInput

# TODO update the JSON string below
json = "{}"
# create an instance of TaxBaseInput from a JSON string
tax_base_input_instance = TaxBaseInput.from_json(json)
# print the JSON string representation of the object
print TaxBaseInput.to_json()

# convert the object into a dict
tax_base_input_dict = tax_base_input_instance.to_dict()
# create an instance of TaxBaseInput from a dict
tax_base_input_form_dict = tax_base_input.from_dict(tax_base_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


