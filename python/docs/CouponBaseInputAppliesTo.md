# CouponBaseInputAppliesTo

Set coupon limitations to plans or specific metrics.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**plan_codes** | **List[str]** | An array of plan codes to which the coupon is applicable. By specifying the plan codes in this field, you can restrict the coupon&#39;s usage to specific plans only. | [optional] 
**billable_metric_codes** | **List[str]** | An array of billable metric codes to which the coupon is applicable. By specifying the billable metric codes in this field, you can restrict the coupon&#39;s usage to specific metrics only. | [optional] 

## Example

```python
from lago_api.models.coupon_base_input_applies_to import CouponBaseInputAppliesTo

# TODO update the JSON string below
json = "{}"
# create an instance of CouponBaseInputAppliesTo from a JSON string
coupon_base_input_applies_to_instance = CouponBaseInputAppliesTo.from_json(json)
# print the JSON string representation of the object
print CouponBaseInputAppliesTo.to_json()

# convert the object into a dict
coupon_base_input_applies_to_dict = coupon_base_input_applies_to_instance.to_dict()
# create an instance of CouponBaseInputAppliesTo from a dict
coupon_base_input_applies_to_form_dict = coupon_base_input_applies_to.from_dict(coupon_base_input_applies_to_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


