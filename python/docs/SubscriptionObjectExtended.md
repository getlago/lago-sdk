# SubscriptionObjectExtended


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lago_id** | **str** | Unique identifier assigned to the subscription within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the subscription’s record within the Lago system | 
**external_id** | **str** | The subscription external unique identifier (provided by your own application). | 
**lago_customer_id** | **str** | Unique identifier assigned to the customer within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the customer&#39;s record within the Lago system | 
**external_customer_id** | **str** | The customer external unique identifier (provided by your own application). | 
**billing_time** | **str** | The billing time for the subscription, which can be set as either &#x60;anniversary&#x60; or &#x60;calendar&#x60;. If not explicitly provided, it will default to &#x60;calendar&#x60;. The billing time determines the timing of recurring billing cycles for the subscription. By specifying &#x60;anniversary&#x60;, the billing cycle will be based on the specific date the subscription started (billed fully), while &#x60;calendar&#x60; sets the billing cycle at the first day of the week/month/year (billed with proration). | 
**name** | **str** | The display name of the subscription on an invoice. This field allows for customization of the subscription&#39;s name for billing purposes, especially useful when a single customer has multiple subscriptions using the same plan. | [optional] 
**plan_code** | **str** | The unique code representing the plan to be attached to the customer. This code must correspond to the &#x60;code&#x60; property of one of the active plans. | 
**status** | **str** | The status of the subscription, which can have the following values: - &#x60;pending&#x60;: a previous subscription has been downgraded, and the current one is awaiting automatic activation at the end of the billing period. - &#x60;active&#x60;: the subscription is currently active and applied to the customer. - &#x60;terminated&#x60;: the subscription is no longer active. - &#x60;canceled&#x60;: the subscription has been stopped before its activation. This can occur when two consecutive downgrades have been applied to a customer or when a subscription with a pending status is terminated. | 
**created_at** | **datetime** | The creation date of the subscription, represented in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). This date provides a timestamp indicating when the subscription was initially created. | 
**canceled_at** | **datetime** | The cancellation date of the subscription. This field is not null when the subscription is &#x60;canceled&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**started_at** | **datetime** | The effective start date of the subscription. This field can be null if the subscription is &#x60;pending&#x60; or &#x60;canceled&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | [optional] 
**ending_at** | **datetime** | The effective end date of the subscription. If this field is set to null, the subscription will automatically renew. This date should be provided in ISO 8601 datetime format, and use Coordinated Universal Time (UTC). | [optional] 
**subscription_at** | **datetime** | The anniversary date and time of the initial subscription. This date serves as the basis for billing subscriptions with &#x60;anniversary&#x60; billing time. The &#x60;anniversary_date&#x60; should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC). | 
**terminated_at** | **datetime** | The termination date of the subscription. This field is not null when the subscription is &#x60;terminated&#x60;. This date should be provided in ISO 8601 datetime format and expressed in Coordinated Universal Time (UTC) | [optional] 
**previous_plan_code** | **str** | The code identifying the previous plan associated with this subscription. | [optional] 
**next_plan_code** | **str** | The code identifying the next plan in the case of a downgrade. | [optional] 
**downgrade_plan_date** | **datetime** | The date when the plan will be downgraded, represented in ISO 8601 date format. | [optional] 
**plan** | [**PlanObject**](PlanObject.md) |  | [optional] 

## Example

```python
from lago_api.models.subscription_object_extended import SubscriptionObjectExtended

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionObjectExtended from a JSON string
subscription_object_extended_instance = SubscriptionObjectExtended.from_json(json)
# print the JSON string representation of the object
print SubscriptionObjectExtended.to_json()

# convert the object into a dict
subscription_object_extended_dict = subscription_object_extended_instance.to_dict()
# create an instance of SubscriptionObjectExtended from a dict
subscription_object_extended_form_dict = subscription_object_extended.from_dict(subscription_object_extended_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


