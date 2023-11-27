# AppliedCouponObjectExtended


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier of the applied coupon, created by Lago. | 
**lago_coupon_id** | **str** | Unique identifier of the coupon, created by Lago. | 
**coupon_code** | **str** | Unique code used to identify the coupon. | 
**coupon_name** | **str** | The name of the coupon. | 
**lago_customer_id** | **str** | Unique identifier of the customer, created by Lago. | 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | 
**status** | **str** | The status of the coupon. Can be either &#x60;active&#x60; or &#x60;terminated&#x60;. | 
**amount_cents** | **int** | The amount of the coupon in cents. This field is required only for coupon with &#x60;fixed_amount&#x60; type. | [optional] 
**amount_cents_remaining** | **int** | The remaining amount in cents for a &#x60;fixed_amount&#x60; coupon with a frequency set to &#x60;once&#x60;. This field indicates the remaining balance or value that can still be utilized from the coupon. | [optional] 
**amount_currency** | [**Currency**](Currency.md) |  | [optional] 
**percentage_rate** | **str** | The percentage rate of the coupon. This field is required only for coupons with a &#x60;percentage&#x60; coupon type. | [optional] 
**frequency** | **str** | The type of frequency for the coupon. It can have three possible values: &#x60;once&#x60;, &#x60;recurring&#x60; or &#x60;forever&#x60;.  - If set to &#x60;once&#x60;, the coupon is applicable only for a single use. - If set to &#x60;recurring&#x60;, the coupon can be used multiple times for recurring billing periods. - If set to &#x60;forever&#x60;, the coupon has unlimited usage and can be applied indefinitely. | 
**frequency_duration** | **int** | Specifies the number of billing periods to which the coupon applies. This field is required only for coupons with a &#x60;recurring&#x60; frequency type | [optional] 
**frequency_duration_remaining** | **int** | The remaining number of billing periods to which the coupon is applicable. This field determines the remaining usage or availability of the coupon based on the remaining billing periods. | [optional] 
**expiration_at** | **datetime** | The date and time after which the coupon will stop applying to customer&#39;s invoices. Once the expiration date is reached, the coupon will no longer be applicable, and any further invoices generated for the customer will not include the coupon discount. | [optional] 
**created_at** | **datetime** | The date and time when the coupon was assigned to a customer. It is expressed in UTC format according to the ISO 8601 datetime standard. | 
**terminated_at** | **datetime** | This field indicates the specific moment when the coupon amount is fully utilized or when the coupon is removed from the customer&#39;s coupon list. It is expressed in UTC format according to the ISO 8601 datetime standard. | [optional] 
**credits** | [**List[CreditObject]**](CreditObject.md) |  | 

## Example

```python
from lago_api.models.applied_coupon_object_extended import AppliedCouponObjectExtended

# TODO update the JSON string below
json = "{}"
# create an instance of AppliedCouponObjectExtended from a JSON string
applied_coupon_object_extended_instance = AppliedCouponObjectExtended.from_json(json)
# print the JSON string representation of the object
print AppliedCouponObjectExtended.to_json()

# convert the object into a dict
applied_coupon_object_extended_dict = applied_coupon_object_extended_instance.to_dict()
# create an instance of AppliedCouponObjectExtended from a dict
applied_coupon_object_extended_form_dict = applied_coupon_object_extended.from_dict(applied_coupon_object_extended_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


