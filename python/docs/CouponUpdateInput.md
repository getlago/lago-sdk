# CouponUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**coupon** | [**CouponBaseInput**](CouponBaseInput.md) |  | 

## Example

```python
from lago_api.models.coupon_update_input import CouponUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CouponUpdateInput from a JSON string
coupon_update_input_instance = CouponUpdateInput.from_json(json)
# print the JSON string representation of the object
print CouponUpdateInput.to_json()

# convert the object into a dict
coupon_update_input_dict = coupon_update_input_instance.to_dict()
# create an instance of CouponUpdateInput from a dict
coupon_update_input_form_dict = coupon_update_input.from_dict(coupon_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


