# FeeUpdateInputFee


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**payment_status** | **str** | The payment status of the fee. Possible values are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; or &#x60;refunded&#x60;. | 

## Example

```python
from lago_api.models.fee_update_input_fee import FeeUpdateInputFee

# TODO update the JSON string below
json = "{}"
# create an instance of FeeUpdateInputFee from a JSON string
fee_update_input_fee_instance = FeeUpdateInputFee.from_json(json)
# print the JSON string representation of the object
print FeeUpdateInputFee.to_json()

# convert the object into a dict
fee_update_input_fee_dict = fee_update_input_fee_instance.to_dict()
# create an instance of FeeUpdateInputFee from a dict
fee_update_input_fee_form_dict = fee_update_input_fee.from_dict(fee_update_input_fee_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


