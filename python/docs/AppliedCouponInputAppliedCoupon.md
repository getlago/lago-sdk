# AppliedCouponInputAppliedCoupon


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | 
**coupon_code** | **str** | Unique code used to identify the coupon. | 
**frequency** | **str** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | [optional] 
**frequency_duration** | **int** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**amount_cents** | **int** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**percentage_rate** | **str** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 

## Example

```python
from lago_api.models.applied_coupon_input_applied_coupon import AppliedCouponInputAppliedCoupon

# TODO update the JSON string below
json = "{}"
# create an instance of AppliedCouponInputAppliedCoupon from a JSON string
applied_coupon_input_applied_coupon_instance = AppliedCouponInputAppliedCoupon.from_json(json)
# print the JSON string representation of the object
print AppliedCouponInputAppliedCoupon.to_json()

# convert the object into a dict
applied_coupon_input_applied_coupon_dict = applied_coupon_input_applied_coupon_instance.to_dict()
# create an instance of AppliedCouponInputAppliedCoupon from a dict
applied_coupon_input_applied_coupon_form_dict = applied_coupon_input_applied_coupon.from_dict(applied_coupon_input_applied_coupon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


