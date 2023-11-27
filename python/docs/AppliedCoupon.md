# AppliedCoupon


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applied_coupon** | [**AppliedCouponObject**](AppliedCouponObject.md) |  | 

## Example

```python
from lago_api.models.applied_coupon import AppliedCoupon

# TODO update the JSON string below
json = "{}"
# create an instance of AppliedCoupon from a JSON string
applied_coupon_instance = AppliedCoupon.from_json(json)
# print the JSON string representation of the object
print AppliedCoupon.to_json()

# convert the object into a dict
applied_coupon_dict = applied_coupon_instance.to_dict()
# create an instance of AppliedCoupon from a dict
applied_coupon_form_dict = applied_coupon.from_dict(applied_coupon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


