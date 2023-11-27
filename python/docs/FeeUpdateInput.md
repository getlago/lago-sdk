# FeeUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fee** | [**FeeUpdateInputFee**](FeeUpdateInputFee.md) |  | 

## Example

```python
from lago_api.models.fee_update_input import FeeUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of FeeUpdateInput from a JSON string
fee_update_input_instance = FeeUpdateInput.from_json(json)
# print the JSON string representation of the object
print FeeUpdateInput.to_json()

# convert the object into a dict
fee_update_input_dict = fee_update_input_instance.to_dict()
# create an instance of FeeUpdateInput from a dict
fee_update_input_form_dict = fee_update_input.from_dict(fee_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


