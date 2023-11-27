# CouponCreateInputCoupon


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the coupon. | 
**code** | **str** | Unique code used to identify the coupon. | 
**description** | **str** | Description of the coupon. | [optional] 
**coupon_type** | **str** | The type of the coupon. It can have two possible values: &#x60;fixed_amount&#x60; or &#x60;percentage&#x60;.  - If set to &#x60;fixed_amount&#x60;, the coupon represents a fixed amount discount. - If set to &#x60;percentage&#x60;, the coupon represents a percentage-based discount. | 
**amount_cents** | **int** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**reusable** | **bool** | Indicates whether the coupon can be reused or not. If set to &#x60;true&#x60;, the coupon is reusable, meaning it can be applied multiple times to the same customer. If set to &#x60;false&#x60;, the coupon can only be used once and is not reusable. If not specified, this field is set to &#x60;true&#x60; by default. | [optional] 
**percentage_rate** | **str** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 
**frequency** | **str** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | 
**frequency_duration** | **int** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**expiration** | **str** | Specifies the type of expiration for the coupon. It can have two possible values: &#x60;time_limit&#x60; or &#x60;no_expiration&#x60;.  - If set to &#x60;time_limit&#x60;, the coupon has an expiration based on a specified time limit. - If set to &#x60;no_expiration&#x60;, the coupon does not have an expiration date and remains valid indefinitely. | [optional] 
**expiration_at** | **datetime** | The expiration date and time of the coupon. This field is required only for coupons with &#x60;expiration&#x60; set to &#x60;time_limit&#x60;. The expiration date and time should be specified in UTC format according to the ISO 8601 datetime standard. It indicates the exact moment when the coupon will expire and is no longer valid. | [optional] 
**applies_to** | [**CouponBaseInputAppliesTo**](CouponBaseInputAppliesTo.md) |  | [optional] 

## Example

```python
from lago_api.models.coupon_create_input_coupon import CouponCreateInputCoupon

# TODO update the JSON string below
json = "{}"
# create an instance of CouponCreateInputCoupon from a JSON string
coupon_create_input_coupon_instance = CouponCreateInputCoupon.from_json(json)
# print the JSON string representation of the object
print CouponCreateInputCoupon.to_json()

# convert the object into a dict
coupon_create_input_coupon_dict = coupon_create_input_coupon_instance.to_dict()
# create an instance of CouponCreateInputCoupon from a dict
coupon_create_input_coupon_form_dict = coupon_create_input_coupon.from_dict(coupon_create_input_coupon_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


