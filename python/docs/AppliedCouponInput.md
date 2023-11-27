# AppliedCouponInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applied_coupon** | [**AppliedCouponInputAppliedCoupon**](AppliedCouponInputAppliedCoupon.md) |  | 

## Example

```python
from lago_api.models.applied_coupon_input import AppliedCouponInput

# TODO update the JSON string below
json = "{}"
# create an instance of AppliedCouponInput from a JSON string
applied_coupon_input_instance = AppliedCouponInput.from_json(json)
# print the JSON string representation of the object
print AppliedCouponInput.to_json()

# convert the object into a dict
applied_coupon_input_dict = applied_coupon_input_instance.to_dict()
# create an instance of AppliedCouponInput from a dict
applied_coupon_input_form_dict = applied_coupon_input.from_dict(applied_coupon_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


