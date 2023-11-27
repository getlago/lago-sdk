# SubscriptionUpdateInputSubscription


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**ending_at** | **datetime** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**subscription_at** | **datetime** | The start date and time of the subscription. This field can only be modified for pending subscriptions that have not yet started. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**plan_overrides** | [**PlanOverridesObject**](PlanOverridesObject.md) |  | [optional] 

## Example

```python
from lago_api.models.subscription_update_input_subscription import SubscriptionUpdateInputSubscription

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionUpdateInputSubscription from a JSON string
subscription_update_input_subscription_instance = SubscriptionUpdateInputSubscription.from_json(json)
# print the JSON string representation of the object
print SubscriptionUpdateInputSubscription.to_json()

# convert the object into a dict
subscription_update_input_subscription_dict = subscription_update_input_subscription_instance.to_dict()
# create an instance of SubscriptionUpdateInputSubscription from a dict
subscription_update_input_subscription_form_dict = subscription_update_input_subscription.from_dict(subscription_update_input_subscription_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


