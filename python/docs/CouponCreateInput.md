# CouponCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**coupon** | [**CouponCreateInputCoupon**](CouponCreateInputCoupon.md) |  | 

## Example

```python
from lago_api.models.coupon_create_input import CouponCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of CouponCreateInput from a JSON string
coupon_create_input_instance = CouponCreateInput.from_json(json)
# print the JSON string representation of the object
print CouponCreateInput.to_json()

# convert the object into a dict
coupon_create_input_dict = coupon_create_input_instance.to_dict()
# create an instance of CouponCreateInput from a dict
coupon_create_input_form_dict = coupon_create_input.from_dict(coupon_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


