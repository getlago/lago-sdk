# Coupon


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**coupon** | [**CouponObject**](CouponObject.md) |  | 

## Example

```python
from lago_api.models.coupon import Coupon

# TODO update the JSON string below
json = "{}"
# create an instance of Coupon from a JSON string
coupon_instance = Coupon.from_json(json)
# print the JSON string representation of the object
print Coupon.to_json()

# convert the object into a dict
coupon_dict = coupon_instance.to_dict()
# create an instance of Coupon from a dict
coupon_form_dict = coupon.from_dict(coupon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


