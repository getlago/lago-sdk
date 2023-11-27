# SubscriptionCreateInputSubscription


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application) | 
**plan_code** | **str** | The unique code representing the plan to be attached to the customer. This code must correspond to the &#x60;code&#x60; property of one of the active plans. | 
**name** | **str** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**external_id** | **str** | The unique external identifier for the subscription. This identifier serves as an idempotency key, ensuring that each subscription is unique. | 
**billing_time** | **str** | The billing time for the subscription, which can be set as either &#x60;anniversary&#x60; or &#x60;calendar&#x60;. If not explicitly provided, it will default to &#x60;calendar&#x60;. The billing time determines the timing of recurring billing cycles for the subscription. By specifying &#x60;anniversary&#x60;, the billing cycle will be based on the specific date the subscription started (billed fully), while &#x60;calendar&#x60; sets the billing cycle at the first day of the week/month/year (billed with proration). | [optional] 
**ending_at** | **datetime** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**subscription_at** | **datetime** | The start date for the subscription, allowing for the creation of subscriptions that can begin in the past or future. Please note that it cannot be used to update the start date of a pending subscription or schedule an upgrade/downgrade. The start_date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**plan_overrides** | [**PlanOverridesObject**](PlanOverridesObject.md) |  | [optional] 

## Example

```python
from lago_api.models.subscription_create_input_subscription import SubscriptionCreateInputSubscription

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionCreateInputSubscription from a JSON string
subscription_create_input_subscription_instance = SubscriptionCreateInputSubscription.from_json(json)
# print the JSON string representation of the object
print SubscriptionCreateInputSubscription.to_json()

# convert the object into a dict
subscription_create_input_subscription_dict = subscription_create_input_subscription_instance.to_dict()
# create an instance of SubscriptionCreateInputSubscription from a dict
subscription_create_input_subscription_form_dict = subscription_create_input_subscription.from_dict(subscription_create_input_subscription_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


